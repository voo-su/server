package repo

import (
	"context"
	"gorm.io/gorm"
	"voo.su/internal/constant"
	"voo.su/internal/domain/entity"
	"voo.su/internal/repository/model"
	"voo.su/pkg/repo"
)

type GroupChatRequest struct {
	repo.Repo[model.GroupChatRequest]
}

func NewGroupChatApply(db *gorm.DB) *GroupChatRequest {
	return &GroupChatRequest{Repo: repo.NewRepo[model.GroupChatRequest](db)}
}

func (g *GroupChatRequest) List(ctx context.Context, groupIds []int) ([]*entity.GroupApplyList, error) {
	fields := []string{
		"group_chat_requests.id",
		"group_chat_requests.group_id",
		"group_chat_requests.user_id",
		"group_chat_requests.created_at",
		"users.avatar",
		"users.username",
	}
	query := g.Repo.Db.WithContext(ctx).Table("group_chat_requests")
	query.Joins("LEFT JOIN users on users.id = group_chat_requests.user_id")
	query.Where("group_chat_requests.status = ?", constant.GroupChatRequestStatusWait)
	query.Order("group_chat_requests.updated_at desc,group_chat_requests.id desc")

	var items []*entity.GroupApplyList
	if err := query.Select(fields).Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}
