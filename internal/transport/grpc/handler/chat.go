package handler

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	chatPb "voo.su/api/grpc/pb"
	"voo.su/internal/config"
	"voo.su/internal/repository/cache"
	"voo.su/internal/usecase"
	"voo.su/pkg/timeutil"
)

type ChatHandler struct {
	chatPb.UnimplementedChatServiceServer
	Conf           *config.Config
	ContactUseCase *usecase.ContactUseCase
	ChatUseCase    *usecase.ChatUseCase
	MessageUseCase usecase.IMessageUseCase
	MessageCache   *cache.MessageCache
	UnreadCache    *cache.UnreadCache
}

func NewChatHandler(
	conf *config.Config,
	contactUseCase *usecase.ContactUseCase,
	chatUseCase *usecase.ChatUseCase,
	messageCache *cache.MessageCache,
	unreadCache *cache.UnreadCache,
) *ChatHandler {
	return &ChatHandler{
		Conf:           conf,
		ContactUseCase: contactUseCase,
		ChatUseCase:    chatUseCase,
		MessageCache:   messageCache,
		UnreadCache:    unreadCache,
	}
}

func (c *ChatHandler) List(ctx context.Context, in *chatPb.GetChatListRequest) (*chatPb.GetChatListResponse, error) {

	// TODO
	uid := 1

	unReads := c.UnreadCache.All(ctx, uid)
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

		if msg, err := c.MessageCache.Get(ctx, item.DialogType, uid, item.ReceiverId); err == nil {
			value.MsgText = msg.Content
			value.UpdatedAt = msg.Datetime
		}
		items = append(items, value)
	}

	return &chatPb.GetChatListResponse{
		Items: items,
	}, nil
}
