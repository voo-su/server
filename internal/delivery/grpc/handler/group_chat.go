package handler

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	groupChatPb "voo.su/api/grpc/pb"
	"voo.su/internal/config"
	"voo.su/internal/constant"
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

func (g *GroupChat) CreateGroupChat(ctx context.Context, in *groupChatPb.CreateGroupChatRequest) (*groupChatPb.CreateGroupChatResponse, error) {
	uid := grpcutil.UserId(ctx)
	gid, err := g.GroupChatUseCase.Create(ctx, &usecase.GroupCreateOpt{
		UserId: uid,
		Name:   in.Name,
	})
	if err != nil {
		return nil, status.Error(codes.Unknown, g.Locale.Localize("chat_creation_failed"))
	}

	return &groupChatPb.CreateGroupChatResponse{
		Id: int64(gid),
	}, nil
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

func (g *GroupChat) AddUserToGroupChat(ctx context.Context, in *groupChatPb.AddUserToGroupChatRequest) (*groupChatPb.AddUserToGroupChatResponse, error) {
	uid := grpcutil.UserId(ctx)

	key := fmt.Sprintf("group-join:%d", in.Id)
	if !g.GroupChatUseCase.RedisLockCacheRepo.Lock(ctx, key, 20) {
		return nil, status.Error(codes.Unknown, g.Locale.Localize("network_error"))
	}

	defer g.GroupChatUseCase.RedisLockCacheRepo.UnLock(ctx, key)

	group, err := g.GroupChatUseCase.GroupChatRepo.FindById(ctx, int(in.Id))
	if err != nil {
		return nil, status.Error(codes.Unknown, g.Locale.Localize("network_error"))
	}

	if group != nil && group.IsDismiss == 1 {
		return nil, status.Error(codes.Unknown, g.Locale.Localize("group_dissolved"))
	}

	var uids []int
	for _, id := range in.UserIds {
		uids = append(uids, int(id))
	}

	if len(uids) == 0 {
		return nil, status.Error(codes.Unknown, g.Locale.Localize("invited_friends_empty"))
	}

	if !g.GroupChatMemberUseCase.MemberRepo.IsMember(ctx, int(in.Id), uid, true) {
		return nil, status.Error(codes.Unknown, g.Locale.Localize("not_group_member_invite"))
	}

	if err := g.GroupChatUseCase.Invite(ctx, &usecase.GroupInviteOpt{
		UserId:    uid,
		GroupId:   int(in.Id),
		MemberIds: uids,
	}); err != nil {
		log.Println(err)
		return nil, status.Error(codes.Unknown, g.Locale.Localize("invite_friends_failed"))
	}

	return &groupChatPb.AddUserToGroupChatResponse{}, nil
}

func (g *GroupChat) RemoveUserFromGroupChat(ctx context.Context, in *groupChatPb.RemoveUserFromGroupChatRequest) (*groupChatPb.RemoveUserFromGroupChatResponse, error) {
	uid := grpcutil.UserId(ctx)
	if !g.GroupChatMemberUseCase.MemberRepo.IsLeader(ctx, int(in.Id), uid) {
		return nil, status.Error(codes.Unknown, g.Locale.Localize("no_permission_for_action"))
	}
	var uids []int
	uids = append(uids, int(in.UserId))

	if err := g.GroupChatUseCase.RemoveMember(ctx, &usecase.GroupRemoveMembersOpt{
		UserId:    uid,
		GroupId:   int(in.Id),
		MemberIds: uids,
	}); err != nil {
		log.Println(err)
		return nil, status.Error(codes.Unknown, g.Locale.Localize("network_error"))
	}

	return &groupChatPb.RemoveUserFromGroupChatResponse{
		Success: true,
	}, nil
}

func (g *GroupChat) LeaveGroupChat(ctx context.Context, in *groupChatPb.LeaveGroupChatRequest) (*groupChatPb.LeaveGroupChatResponse, error) {
	uid := grpcutil.UserId(ctx)
	if err := g.GroupChatUseCase.Secede(ctx, int(in.Id), uid); err != nil {
		log.Println(err)
		return nil, status.Error(codes.Unknown, g.Locale.Localize("network_error"))
	}

	sid := g.ChatUseCase.ChatRepo.FindBySessionId(uid, int(in.Id), constant.ChatGroupMode)
	if err := g.ChatUseCase.Delete(ctx, uid, sid); err != nil {
		log.Println(err)
	}

	return &groupChatPb.LeaveGroupChatResponse{
		Success: true,
	}, nil
}
