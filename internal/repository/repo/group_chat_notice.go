package repo

import (
	"context"
	"gorm.io/gorm"
	"voo.su/internal/repository/model"
	"voo.su/pkg/core"
)

type GroupChatNotice struct {
	core.Repo[model.GroupChatNotice]
}

func NewGroupChatNotice(db *gorm.DB) *GroupChatNotice {
	return &GroupChatNotice{Repo: core.NewRepo[model.GroupChatNotice](db)}
}

func (g *GroupChatNotice) GetListAll(ctx context.Context, groupId int) ([]*model.SearchNoticeItem, error) {
	fields := []string{
		"group_chat_notice.id",
		"group_chat_notice.creator_id",
		"group_chat_notice.title",
		"group_chat_notice.content",
		"group_chat_notice.is_top",
		"group_chat_notice.is_confirm",
		"group_chat_notice.confirm_users",
		"group_chat_notice.created_at",
		"group_chat_notice.updated_at",
		"users.avatar",
		"users.username",
	}
	query := g.Repo.Db.WithContext(ctx).Table("group_chat_notice")
	query.Joins("LEFT JOIN users on users.id = group_chat_notice.creator_id")
	query.Where("group_chat_notice.group_id = ? and group_chat_notice.is_delete = ?", groupId, 0)
	query.Order("group_chat_notice.is_top desc")
	query.Order("group_chat_notice.created_at desc")

	var items []*model.SearchNoticeItem
	if err := query.Select(fields).Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (g *GroupChatNotice) GetLatestNotice(ctx context.Context, groupId int) (*model.GroupChatNotice, error) {
	var info model.GroupChatNotice
	err := g.Repo.Db.WithContext(ctx).Last(&info, "group_id = ? and is_delete = ?", groupId, 0).Error
	if err != nil {
		return nil, err
	}

	return &info, nil
}
