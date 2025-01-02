// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package repository

import (
	"context"
	"gorm.io/gorm"
	"voo.su/internal/infrastructure/postgres/model"
	"voo.su/pkg/repo"
)

type GroupChatRepository struct {
	repo.Repo[model.GroupChat]
}

func NewGroupChatRepository(db *gorm.DB) *GroupChatRepository {
	return &GroupChatRepository{Repo: repo.NewRepo[model.GroupChat](db)}
}

type SearchOvertListOpt struct {
	Name   string
	UserId int
	Page   int
	Size   int
}

func (g *GroupChatRepository) SearchOvertList(ctx context.Context, opt *SearchOvertListOpt) ([]*model.GroupChat, error) {
	return g.Repo.FindAll(ctx, func(db *gorm.DB) {
		if opt.Name != "" {
			db.Where("group_name like ?", "%"+opt.Name+"%")
		}
		db.Where("is_overt = ?", 1).
			Where("id NOT IN (?)", g.Repo.Db.Select("group_id").
				Where("user_id = ? AND is_quit= ?", opt.UserId, 0).
				Table("group_chat_members"),
			).
			Where("is_dismiss = ?", 0).
			Order("created_at desc").
			Offset((opt.Page - 1) * opt.Size).
			Limit(opt.Size)
	})
}
