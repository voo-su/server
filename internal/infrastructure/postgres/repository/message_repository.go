package repository

import (
	"gorm.io/gorm"
	"voo.su/internal/infrastructure/postgres/model"
	"voo.su/pkg/repo"
)

type MessageRepository struct {
	repo.Repo[model.Message]
}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{Repo: repo.NewRepo[model.Message](db)}
}
