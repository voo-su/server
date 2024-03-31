package service

import (
	"context"
	"gorm.io/gorm"
	groupChat "voo.su/internal/repository/model"
	"voo.su/internal/repository/repo"
)

var _ IGroupChatMemberService = (*GroupChatMemberService)(nil)

type IGroupChatMemberService interface {
	Handover(ctx context.Context, groupId int, userId int, memberId int) error
	SetLeaderStatus(ctx context.Context, groupId int, userId int, leader int) error
	SetMuteStatus(ctx context.Context, groupId int, userId int, status int) error
}

type GroupChatMemberService struct {
	*repo.Source
	member *repo.GroupChatMember
}

func NewGroupMemberService(source *repo.Source, repo *repo.GroupChatMember) *GroupChatMemberService {
	return &GroupChatMemberService{Source: source, member: repo}
}

func (g *GroupChatMemberService) Handover(ctx context.Context, groupId int, userId int, memberId int) error {
	return g.Source.Db().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&groupChat.GroupChatMember{}).
			Where("group_id = ? and user_id = ? and leader = 2", groupId, userId).
			Update("leader", 0).Error
		if err != nil {
			return err
		}

		err = tx.Model(&groupChat.GroupChatMember{}).
			Where("group_id = ? and user_id = ?", groupId, memberId).
			Update("leader", 2).Error
		if err != nil {
			return err
		}

		return nil
	})
}

func (g *GroupChatMemberService) SetLeaderStatus(ctx context.Context, groupId int, userId int, leader int) error {
	return g.member.Model(ctx).
		Where("group_id = ? and user_id = ?", groupId, userId).
		UpdateColumn("leader", leader).
		Error
}

func (g *GroupChatMemberService) SetMuteStatus(ctx context.Context, groupId int, userId int, status int) error {
	return g.member.Model(ctx).
		Where("group_id = ? and user_id = ?", groupId, userId).
		UpdateColumn("is_mute", status).
		Error
}
