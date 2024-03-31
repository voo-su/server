package repo

import (
	"context"
	"gorm.io/gorm"
	"voo.su/internal/repository/model"
	"voo.su/pkg/core"
)

type GroupChat struct {
	core.Repo[model.GroupChat]
}

func NewGroupChat(db *gorm.DB) *GroupChat {
	return &GroupChat{Repo: core.NewRepo[model.GroupChat](db)}
}

type SearchOvertListOpt struct {
	Name   string
	UserId int
	Page   int
	Size   int
}

func (g *GroupChat) SearchOvertList(ctx context.Context, opt *SearchOvertListOpt) ([]*model.GroupChat, error) {
	return g.Repo.FindAll(ctx, func(db *gorm.DB) {
		if opt.Name != "" {
			db.Where("group_name like ?", "%"+opt.Name+"%")
		}
		db.Where("is_overt = ?", 1)
		db.Where("id NOT IN (?)", g.Repo.Db.Select("group_id").
			Where("user_id = ? and is_quit= ?", opt.UserId, 0).
			Table("group_chat_members"),
		)
		db.Where("is_dismiss = 0").
			Order("created_at desc").
			Offset((opt.Page - 1) * opt.Size).
			Limit(opt.Size)
	})
}
