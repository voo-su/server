package repository

import (
	"context"
	"gorm.io/gorm"
	"voo.su/internal/infrastructure/postgres/model"
	"voo.su/pkg/repo"
)

type BotRepository struct {
	repo.Repo[model.Bot]
}

func NewBotRepository(db *gorm.DB) *BotRepository {
	return &BotRepository{Repo: repo.NewRepo[model.Bot](db)}
}

func (b *BotRepository) GetLoginBot(ctx context.Context) (*model.Bot, error) {
	return b.Repo.FindByWhere(ctx, "bot_type = ?", 1)
}

func (b *BotRepository) GetByUserId(ctx context.Context, userId int) (*model.Bot, error) {
	return b.Repo.FindByWhere(ctx, "user_id = ? AND bot_type = ?", userId, 1)
}
