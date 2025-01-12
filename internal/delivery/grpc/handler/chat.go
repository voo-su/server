package handler

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	chatPb "voo.su/api/grpc/gen/go/pb"
	"voo.su/internal/config"
	"voo.su/internal/usecase"
	"voo.su/pkg/grpcutil"
	"voo.su/pkg/locale"
	"voo.su/pkg/timeutil"
)

type Chat struct {
	chatPb.UnimplementedChatServiceServer
	Conf           *config.Config
	Locale         locale.ILocale
	ContactUseCase *usecase.ContactUseCase
	ChatUseCase    *usecase.ChatUseCase
	MessageUseCase usecase.IMessageUseCase
}

func NewChatHandler(
	conf *config.Config,
	locale locale.ILocale,
	contactUseCase *usecase.ContactUseCase,
	chatUseCase *usecase.ChatUseCase,
) *Chat {
	return &Chat{
		Conf:           conf,
		Locale:         locale,
		ContactUseCase: contactUseCase,
		ChatUseCase:    chatUseCase,
	}
}

func (c *Chat) List(ctx context.Context, in *chatPb.GetChatListRequest) (*chatPb.GetChatListResponse, error) {
	uid := grpcutil.UserId(ctx)

	unReads := c.ChatUseCase.UnreadCacheRepo.All(ctx, uid)
	if len(unReads) > 0 {
		c.ChatUseCase.BatchAddList(ctx, uid, unReads)
	}

	data, err := c.ChatUseCase.List(ctx, 1)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	items := make([]*chatPb.ChatItem, 0)
	for _, item := range data {
		value := &chatPb.ChatItem{
			Id:        int32(item.Id),
			ChatType:  int32(item.DialogType),
			Avatar:    item.UserAvatar,
			MsgText:   "",
			UpdatedAt: timeutil.FormatDatetime(item.UpdatedAt),
		}

		if num, ok := unReads[fmt.Sprintf("%d_%d", item.DialogType, item.ReceiverId)]; ok {
			value.UnreadNum = int32(num)
		}

		if item.DialogType == 1 {
			value.Username = item.Username
			value.Avatar = item.UserAvatar
			value.Name = item.Name
			value.Surname = item.Surname
		} else {
			value.Name = item.GroupName
			value.Avatar = item.GroupAvatar
		}

		if msg, err := c.ChatUseCase.MessageCacheRepo.Get(ctx, item.DialogType, uid, item.ReceiverId); err == nil {
			value.MsgText = msg.Content
			value.UpdatedAt = msg.Datetime
		}
		items = append(items, value)
	}

	return &chatPb.GetChatListResponse{
		Items: items,
	}, nil
}
