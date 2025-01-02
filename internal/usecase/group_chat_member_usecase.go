// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package usecase

import (
	"context"
	"gorm.io/gorm"
	"voo.su/internal/infrastructure"
	postgresModel "voo.su/internal/infrastructure/postgres/model"
	postgresRepo "voo.su/internal/infrastructure/postgres/repository"
	"voo.su/pkg/locale"
)

type IGroupChatMemberUseCase interface {
	Handover(ctx context.Context, groupId int, userId int, memberId int) error

	SetLeaderStatus(ctx context.Context, groupId int, userId int, leader int) error

	SetMuteStatus(ctx context.Context, groupId int, userId int, status int) error
}

var _ IGroupChatMemberUseCase = (*GroupChatMemberUseCase)(nil)

type GroupChatMemberUseCase struct {
	Locale     locale.ILocale
	Source     *infrastructure.Source
	MemberRepo *postgresRepo.GroupChatMemberRepository
}

func NewGroupMemberUseCase(
	locale locale.ILocale,
	source *infrastructure.Source,
	memberRepo *postgresRepo.GroupChatMemberRepository,
) *GroupChatMemberUseCase {
	return &GroupChatMemberUseCase{
		Locale:     locale,
		Source:     source,
		MemberRepo: memberRepo,
	}
}

func (g *GroupChatMemberUseCase) Handover(ctx context.Context, groupId int, userId int, memberId int) error {
	return g.Source.Postgres().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&postgresModel.GroupChatMember{}).
			Where("group_id = ? AND user_id = ? AND leader = ?", groupId, userId, 2).
			Update("leader", 0).
			Error; err != nil {
			return err
		}

		if err := tx.Model(&postgresModel.GroupChatMember{}).
			Where("group_id = ? AND user_id = ?", groupId, memberId).
			Update("leader", 2).
			Error; err != nil {
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
