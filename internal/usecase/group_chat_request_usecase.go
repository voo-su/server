package usecase

import (
	"context"
	"errors"
	"voo.su/internal/infrastructure"
	postgresModel "voo.su/internal/infrastructure/postgres/model"
	postgresRepo "voo.su/internal/infrastructure/postgres/repository"
	redisRepo "voo.su/internal/infrastructure/redis/repository"
	"voo.su/pkg/locale"
)

type GroupChatRequestUseCase struct {
	Locale                locale.ILocale
	Source                *infrastructure.Source
	GroupChatRequestRepo  *postgresRepo.GroupChatRequestRepository
	GroupRequestCacheRepo *redisRepo.GroupChatRequestCacheRepository
}

func NewGroupRequestUseCase(
	locale locale.ILocale,
	source *infrastructure.Source,
	groupChatRequestRepo *postgresRepo.GroupChatRequestRepository,
	groupRequestCacheRepo *redisRepo.GroupChatRequestCacheRepository,
) *GroupChatRequestUseCase {
	return &GroupChatRequestUseCase{
		Locale:                locale,
		Source:                source,
		GroupChatRequestRepo:  groupChatRequestRepo,
		GroupRequestCacheRepo: groupRequestCacheRepo,
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

func (g *GroupChatRequestUseCase) Delete(ctx context.Context, id, userId int) error {
	if !g.Auth(ctx, id, userId) {
		return errors.New(g.Locale.Localize("authentication_failed"))
	}

	return g.Source.Postgres().
		WithContext(ctx).
		Delete(&postgresModel.GroupChatRequest{}, "id = ?", id).
		Error
}
