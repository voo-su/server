package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	chatPb "voo.su/api/grpc/pb"
	"voo.su/internal/config"
	"voo.su/internal/service"
)

type ChatHandler struct {
	chatPb.UnimplementedChatServiceServer
	Conf               *config.Config
	ContactService     *service.ContactService
	ChatService        *service.ChatService
	MessageSendService service.MessageSendService
}

func NewChatHandler(
	conf *config.Config,
	contactService *service.ContactService,
	chatService *service.ChatService,
	messageSendService service.MessageSendService,
) *ChatHandler {
	return &ChatHandler{
		Conf:               conf,
		ContactService:     contactService,
		ChatService:        chatService,
		MessageSendService: messageSendService,
	}
}

func (c *ChatHandler) List(ctx context.Context, in *chatPb.ChatListRequest) (*chatPb.ChatListResponse, error) {
	data, err := c.ChatService.List(ctx, 1)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}
	items := make([]*chatPb.ChatItem, 0)
	for _, item := range data {
		value := &chatPb.ChatItem{
			Id:       int32(item.Id),
			ChatType: int32(item.DialogType),
		}

		if item.DialogType == 1 {
			value.Username = item.Username
			value.Name = item.Name
			value.Surname = item.Surname
		} else {
			value.Name = item.GroupName
		}

		items = append(items, value)
	}

	return &chatPb.ChatListResponse{
		Items: items,
	}, nil
}
