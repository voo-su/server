package usecase

import (
	"context"
	"errors"
	"voo.su/internal/repository/model"
	"voo.su/internal/repository/repo"
)

type GroupChatRequestUseCase struct {
	*repo.Source
	Repo *repo.GroupChatRequest
}

func NewGroupRequestUseCase(
	source *repo.Source,
	repo *repo.GroupChatRequest,
) *GroupChatRequestUseCase {
	return &GroupChatRequestUseCase{
		Source: source,
		Repo:   repo,
	}
}

func (s *GroupChatRequestUseCase) Auth(ctx context.Context, id, userId int) bool {
	info, err := s.Repo.FindById(ctx, id)
	if err != nil {
		return false
	}

	var member model.GroupChatMember
	err = s.Source.Db().Debug().WithContext(ctx).
		Select("id").
		First(&member, "group_id = ? and user_id = ? and leader in (1,2) and is_quit = 0", info.GroupId, userId).Error

	return err == nil && member.Id > 0
}

func (s *GroupChatRequestUseCase) Insert(ctx context.Context, groupId, userId int /*, remark string*/) error {
	return s.Repo.Create(ctx, &model.GroupChatRequest{
		GroupId: groupId,
		UserId:  userId,
		//Remark:  remark,
	})
}

func (s *GroupChatRequestUseCase) Delete(ctx context.Context, id, userId int) error {
	if !s.Auth(ctx, id, userId) {
		return errors.New("не удалось выполнить аутентификацию")
	}

	return s.Source.Db().WithContext(ctx).Delete(&model.GroupChatRequest{}, "id = ?", id).Error
}
