package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	groupChatPb "voo.su/api/grpc/gen/go/pb"
	"voo.su/internal/config"
	"voo.su/internal/usecase"
	"voo.su/pkg/grpcutil"
	"voo.su/pkg/locale"
)

type GroupChat struct {
	groupChatPb.UnimplementedGroupChatServiceServer
	Conf                   *config.Config
	Locale                 locale.ILocale
	ContactUseCase         *usecase.ContactUseCase
	ChatUseCase            *usecase.ChatUseCase
	MessageUseCase         usecase.IMessageUseCase
	GroupChatUseCase       *usecase.GroupChatUseCase
	GroupChatMemberUseCase *usecase.GroupChatMemberUseCase
}

func NewGroupChatHandler(
	conf *config.Config,
	locale locale.ILocale,
	contactUseCase *usecase.ContactUseCase,
	chatUseCase *usecase.ChatUseCase,
	groupChatUseCase *usecase.GroupChatUseCase,
	groupChatMemberUseCase *usecase.GroupChatMemberUseCase,
) *GroupChat {
	return &GroupChat{
		Conf:                   conf,
		Locale:                 locale,
		ContactUseCase:         contactUseCase,
		ChatUseCase:            chatUseCase,
		GroupChatUseCase:       groupChatUseCase,
		GroupChatMemberUseCase: groupChatMemberUseCase,
	}
}

func (g *GroupChat) GetGroupChat(ctx context.Context, in *groupChatPb.GetGroupChatRequest) (*groupChatPb.GetGroupChatResponse, error) {
	uid := grpcutil.UserId(ctx)
	group, err := g.GroupChatUseCase.GroupChatRepo.FindById(ctx, int(in.Id))
	if err != nil {
		return nil, status.Error(codes.Unknown, g.Locale.Localize("network_error"))
	}

	if group != nil && group.IsDismiss == 1 {
		return nil, status.Error(codes.Unknown, g.Locale.Localize("network_error"))
	}

	if !g.GroupChatMemberUseCase.MemberRepo.IsMember(ctx, int(in.Id), uid, false) {
		return nil, status.Error(codes.Unknown, g.Locale.Localize("not_member_cannot_view_participants"))
	}

	countMembers := g.GroupChatMemberUseCase.MemberRepo.CountMemberTotal(ctx, int(in.Id))
	return &groupChatPb.GetGroupChatResponse{
		Id:      int64(group.Id),
		Name:    group.Name,
		Avatar:  group.Avatar,
		Members: countMembers,
	}, nil
}

func (g *GroupChat) GetMembers(ctx context.Context, in *groupChatPb.GetMembersRequest) (*groupChatPb.GetMembersResponse, error) {
	uid := grpcutil.UserId(ctx)
	group, err := g.GroupChatUseCase.GroupChatRepo.FindById(ctx, int(in.Id))
	if err != nil {
		return nil, status.Error(codes.Unknown, g.Locale.Localize("network_error"))
	}

	if group != nil && group.IsDismiss == 1 {
		return nil, status.Error(codes.Unknown, g.Locale.Localize("network_error"))
	}

	if !g.GroupChatMemberUseCase.MemberRepo.IsMember(ctx, int(in.Id), uid, false) {
		return nil, status.Error(codes.Unknown, g.Locale.Localize("not_member_cannot_view_participants"))
	}

	list := g.GroupChatMemberUseCase.MemberRepo.GetMembers(ctx, int(in.Id))

	items := make([]*groupChatPb.MemberItem, 0)
	for _, item := range list {
		items = append(items, &groupChatPb.MemberItem{
			Id:       int64(item.UserId),
			Username: item.Username,
			Avatar:   item.Avatar,
			// TODO
		})
	}

	return &groupChatPb.GetMembersResponse{
		Items: items,
	}, nil
}
