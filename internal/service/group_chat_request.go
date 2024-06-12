package service

import (
	"context"
	"errors"
	"voo.su/internal/repository/model"
	"voo.su/internal/repository/repo"
)

type GroupChatRequestService struct {
	*repo.Source
	Repo *repo.GroupChatRequest
}

func NewGroupRequestService(
	source *repo.Source,
	repo *repo.GroupChatRequest,
) *GroupChatRequestService {
	return &GroupChatRequestService{
		Source: source,
		Repo:   repo,
	}
}

func (s *GroupChatRequestService) Auth(ctx context.Context, id, userId int) bool {
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

func (s *GroupChatRequestService) Insert(ctx context.Context, groupId, userId int /*, remark string*/) error {
	return s.Repo.Create(ctx, &model.GroupChatRequest{
		GroupId: groupId,
		UserId:  userId,
		//Remark:  remark,
	})
}

func (s *GroupChatRequestService) Delete(ctx context.Context, id, userId int) error {
	if !s.Auth(ctx, id, userId) {
		return errors.New("не удалось выполнить аутентификацию")
	}

	return s.Source.Db().WithContext(ctx).Delete(&model.GroupChatRequest{}, "id = ?", id).Error
}
