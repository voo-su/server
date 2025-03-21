package chat

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
	"sync"
	chatPb "voo.su/api/grpc/pb"
	commonPb "voo.su/api/grpc/pb/common"
	"voo.su/internal/config"
	"voo.su/internal/constant"
	"voo.su/internal/usecase"
	"voo.su/pkg/grpcutil"
	"voo.su/pkg/locale"
	"voo.su/pkg/nats"
	"voo.su/pkg/timeutil"
)

type Chat struct {
	chatPb.UnimplementedChatServiceServer
	Conf           *config.Config
	Locale         locale.ILocale
	ContactUseCase *usecase.ContactUseCase
	ChatUseCase    *usecase.ChatUseCase
	MessageUseCase usecase.IMessageUseCase
	mu             sync.Mutex
	userChannels   map[int64]chan *chatPb.Update
	Nats           nats.INatsClient
	UploadUseCase  *usecase.UploadUseCase
}

func NewChatHandler(
	conf *config.Config,
	locale locale.ILocale,
	contactUseCase *usecase.ContactUseCase,
	chatUseCase *usecase.ChatUseCase,
	messageUseCase *usecase.MessageUseCase,
	nats nats.INatsClient,
	uploadUseCase *usecase.UploadUseCase,
) *Chat {
	return &Chat{
		Conf:           conf,
		Locale:         locale,
		ContactUseCase: contactUseCase,
		ChatUseCase:    chatUseCase,
		MessageUseCase: messageUseCase,
		userChannels:   make(map[int64]chan *chatPb.Update),
		Nats:           nats,
		UploadUseCase:  uploadUseCase,
	}
}

func (c *Chat) GetChats(ctx context.Context, in *chatPb.GetChatsRequest) (*chatPb.GetChatsResponse, error) {
	uid := grpcutil.UserId(ctx)
	unReads := c.ChatUseCase.UnreadCacheRepo.All(ctx, uid)
	if len(unReads) > 0 {
		c.ChatUseCase.BatchAddList(ctx, uid, unReads)
	}

	data, err := c.ChatUseCase.List(ctx, uid)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}
	friends := make([]int, 0)
	for _, item := range data {
		if item.ChatType == 1 {
			friends = append(friends, item.ReceiverId)
		}
	}

	items := make([]*chatPb.ChatItem, 0)
	for _, item := range data {
		value := &chatPb.ChatItem{
			Id: int64(item.Id),
			Receiver: &chatPb.Receiver{
				ChatType:   int32(item.ChatType),
				ReceiverId: int64(item.ReceiverId),
			},
			Avatar:  item.UserAvatar,
			MsgText: "",
			NotifySettings: &commonPb.EntityNotifySettings{
				MuteUntil:    item.NotifyMuteUntil,
				ShowPreviews: item.NotifyShowPreviews,
				Silent:       item.NotifySilent,
			},
			UpdatedAt: timeutil.FormatDatetime(item.UpdatedAt),
			IsDisturb: item.IsDisturb == 1,
			IsBot:     item.IsBot == 1,
		}

		if num, ok := unReads[fmt.Sprintf("%d_%d", item.ChatType, item.ReceiverId)]; ok {
			value.UnreadCount = int64(num)
		}

		if item.ChatType == 1 {
			value.Username = item.Username
			value.Avatar = item.UserAvatar
			value.Name = item.Name
			value.Surname = item.Surname
			value.IsOnline = c.ChatUseCase.ClientCacheRepo.IsOnline(ctx, constant.ImChannelChat, strconv.Itoa(int(value.Receiver.ReceiverId)))
		} else {
			value.Name = item.GroupName
			value.Avatar = item.GroupAvatar
		}

		if msg, err := c.ChatUseCase.MessageCacheRepo.Get(ctx, item.ChatType, uid, item.ReceiverId); err == nil {
			value.MsgText = msg.Content
			value.UpdatedAt = msg.Datetime
		}
		items = append(items, value)
	}

	return &chatPb.GetChatsResponse{
		Items: items,
	}, nil
}
