package usecase

import (
	"context"
	"gorm.io/gorm"
	groupChat "voo.su/internal/repository/model"
	"voo.su/internal/repository/repo"
)

type IGroupChatMemberUseCase interface {
	Handover(ctx context.Context, groupId int, userId int, memberId int) error
	SetLeaderStatus(ctx context.Context, groupId int, userId int, leader int) error
	SetMuteStatus(ctx context.Context, groupId int, userId int, status int) error
}

var _ IGroupChatMemberUseCase = (*GroupChatMemberUseCase)(nil)

type GroupChatMemberUseCase struct {
	*repo.Source
	MemberRepo *repo.GroupChatMember
}

func NewGroupMemberUseCase(
	source *repo.Source,
	memberRepo *repo.GroupChatMember,
) *GroupChatMemberUseCase {
	return &GroupChatMemberUseCase{
		Source:     source,
		MemberRepo: memberRepo,
	}
}

func (g *GroupChatMemberUseCase) Handover(ctx context.Context, groupId int, userId int, memberId int) error {
	return g.Source.Db().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&groupChat.GroupChatMember{}).
			Where("group_id = ? AND user_id = ? AND leader = 2", groupId, userId).
			Update("leader", 0).Error
		if err != nil {
			return err
		}

		err = tx.Model(&groupChat.GroupChatMember{}).
			Where("group_id = ? AND user_id = ?", groupId, memberId).
			Update("leader", 2).Error
		if err != nil {
			return err
		}

		return nil
	})
}

func (g *GroupChatMemberUseCase) SetLeaderStatus(ctx context.Context, groupId int, userId int, leader int) error {
	return g.MemberRepo.Model(ctx).
		Where("group_id = ? AND user_id = ?", groupId, userId).
		UpdateColumn("leader", leader).
		Error
}

func (g *GroupChatMemberUseCase) SetMuteStatus(ctx context.Context, groupId int, userId int, status int) error {
	return g.MemberRepo.Model(ctx).
		Where("group_id = ? AND user_id = ?", groupId, userId).
		UpdateColumn("is_mute", status).
		Error
}
