package repo

import (
	"context"
	"gorm.io/gorm"
	"voo.su/internal/repository/model"
	"voo.su/pkg/core"
)

type Bot struct {
	core.Repo[model.Bot]
}

func NewBot(db *gorm.DB) *Bot {
	return &Bot{Repo: core.NewRepo[model.Bot](db)}
}

func (r *Bot) GetLoginBot(ctx context.Context) (*model.Bot, error) {
	return r.Repo.FindByWhere(ctx, "bot_type = ?", 1)
}

func (r *Bot) GetByUserId(ctx context.Context, userId int) (*model.Bot, error) {
	return r.Repo.FindByWhere(ctx, "user_id = ? AND bot_type = ?", userId, 1)
}
