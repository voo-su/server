package usecase

import (
	"context"
	"errors"
	"voo.su/internal/repository"
	"voo.su/internal/repository/model"
	"voo.su/internal/repository/repo"
)

type GroupChatRequestUseCase struct {
	*repository.Source
	GroupChatRequestRepo *repo.GroupChatRequest
}

func NewGroupRequestUseCase(
	source *repository.Source,
	groupChatRequestRepo *repo.GroupChatRequest,
) *GroupChatRequestUseCase {
	return &GroupChatRequestUseCase{
		Source:               source,
		GroupChatRequestRepo: groupChatRequestRepo,
	}
}

func (g *GroupChatRequestUseCase) Auth(ctx context.Context, id, userId int) bool {
	info, err := g.GroupChatRequestRepo.FindById(ctx, id)
	if err != nil {
		return false
	}

	var member model.GroupChatMember
	err = g.Source.Db().
		Debug().
		WithContext(ctx).
		Select("id").
		First(&member, "group_id = ? AND user_id = ? AND leader in (1,2) AND is_quit = 0", info.GroupId, userId).
		Error

	return err == nil && member.Id > 0
}

func (g *GroupChatRequestUseCase) Insert(ctx context.Context, groupId, userId int /*, remark string*/) error {
	return g.GroupChatRequestRepo.Create(ctx, &model.GroupChatRequest{
		GroupId: groupId,
		UserId:  userId,
		//Remark:  remark,
	})
}

func (g *GroupChatRequestUseCase) Delete(ctx context.Context, id, userId int) error {
	if !g.Auth(ctx, id, userId) {
		return errors.New("не удалось выполнить аутентификацию")
	}

	return g.Source.Db().
		WithContext(ctx).
		Delete(&model.GroupChatRequest{}, "id = ?", id).
		Error
}
