package handler

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	chatPb "voo.su/api/grpc/pb"
	"voo.su/internal/config"
	"voo.su/internal/repository/cache"
	"voo.su/internal/service"
	"voo.su/pkg/timeutil"
)

type ChatHandler struct {
	chatPb.UnimplementedChatServiceServer
	Conf               *config.Config
	ContactService     *service.ContactService
	ChatService        *service.ChatService
	MessageSendService service.MessageSendService
	MessageStorage     *cache.MessageStorage
	UnreadStorage      *cache.UnreadStorage
}

func NewChatHandler(
	conf *config.Config,
	contactService *service.ContactService,
	chatService *service.ChatService,
	messageSendService service.MessageSendService,
	messageStorage *cache.MessageStorage,
	unreadStorage *cache.UnreadStorage,
) *ChatHandler {
	return &ChatHandler{
		Conf:               conf,
		ContactService:     contactService,
		ChatService:        chatService,
		MessageSendService: messageSendService,
		MessageStorage:     messageStorage,
		UnreadStorage:      unreadStorage,
	}
}

func (c *ChatHandler) List(ctx context.Context, in *chatPb.GetChatListRequest) (*chatPb.GetChatListResponse, error) {

	// TODO
	uid := 1

	unReads := c.UnreadStorage.All(ctx, uid)
	if len(unReads) > 0 {
		c.ChatService.BatchAddList(ctx, uid, unReads)
	}

	data, err := c.ChatService.List(ctx, 1)
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

		if msg, err := c.MessageStorage.Get(ctx, item.DialogType, uid, item.ReceiverId); err == nil {
			value.MsgText = msg.Content
			value.UpdatedAt = msg.Datetime
		}
		items = append(items, value)
	}

	return &chatPb.GetChatListResponse{
		Items: items,
	}, nil
}
