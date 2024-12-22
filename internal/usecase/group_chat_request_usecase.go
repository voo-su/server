package usecase

import (
	"context"
	"errors"
	"voo.su/internal/infrastructure"
	postgresModel "voo.su/internal/infrastructure/postgres/model"
	postgresRepo "voo.su/internal/infrastructure/postgres/repository"
	"voo.su/pkg/locale"
)

type GroupChatRequestUseCase struct {
	Locale               locale.ILocale
	Source               *infrastructure.Source
	GroupChatRequestRepo *postgresRepo.GroupChatRequestRepository
}

func NewGroupRequestUseCase(
	locale locale.ILocale,
	source *infrastructure.Source,
	groupChatRequestRepo *postgresRepo.GroupChatRequestRepository,
) *GroupChatRequestUseCase {
	return &GroupChatRequestUseCase{
		Locale:               locale,
		Source:               source,
		GroupChatRequestRepo: groupChatRequestRepo,
	}
}

func (g *GroupChatRequestUseCase) Auth(ctx context.Context, id, userId int) bool {
	info, err := g.GroupChatRequestRepo.FindById(ctx, id)
	if err != nil {
		return false
	}

	var member postgresModel.GroupChatMember
	err = g.Source.Postgres().
		Debug().
		WithContext(ctx).
		Select("id").
		First(&member, "group_id = ? AND user_id = ? AND leader IN (1,2) AND is_quit = 0", info.GroupId, userId).
		Error

	return err == nil && member.Id > 0
}

func (g *GroupChatRequestUseCase) Insert(ctx context.Context, groupId, userId int /*, remark string*/) error {
	return g.GroupChatRequestRepo.Create(ctx, &postgresModel.GroupChatRequest{
		GroupId: groupId,
		UserId:  userId,
		//Remark:  remark,
	})
}

func (g *GroupChatRequestUseCase) Delete(ctx context.Context, id, userId int) error {
	if !g.Auth(ctx, id, userId) {
		return errors.New(g.Locale.Localize("authentication_failed"))
	}

	return g.Source.Postgres().
		WithContext(ctx).
		Delete(&postgresModel.GroupChatRequest{}, "id = ?", id).
		Error
}
