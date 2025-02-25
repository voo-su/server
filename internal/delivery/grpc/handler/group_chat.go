package handler

import (
	"context"
	groupChatPb "voo.su/api/grpc/gen/go/pb"
	"voo.su/internal/config"
	"voo.su/internal/usecase"
	"voo.su/pkg/locale"
)

type GroupChat struct {
	groupChatPb.UnimplementedGroupChatServiceServer
	Conf           *config.Config
	Locale         locale.ILocale
	ContactUseCase *usecase.ContactUseCase
	ChatUseCase    *usecase.ChatUseCase
	MessageUseCase usecase.IMessageUseCase
}

func NewGroupChatHandler(
	conf *config.Config,
	locale locale.ILocale,
	contactUseCase *usecase.ContactUseCase,
	chatUseCase *usecase.ChatUseCase,
) *GroupChat {
	return &GroupChat{
		Conf:           conf,
		Locale:         locale,
		ContactUseCase: contactUseCase,
		ChatUseCase:    chatUseCase,
	}
}

func (g *GroupChat) GetGroupChat(ctx context.Context, in *groupChatPb.GetGroupChatRequest) (*groupChatPb.GetGroupChatResponse, error) {
	// TODO
	return &groupChatPb.GetGroupChatResponse{}, nil
}

func (g *GroupChat) GetMembers(ctx context.Context, in *groupChatPb.GetMembersRequest) (*groupChatPb.GetMembersResponse, error) {
	// TODO
	return &groupChatPb.GetMembersResponse{}, nil
}
