package repository

import (
	"gorm.io/gorm"
	"voo.su/internal/infrastructure/postgres/model"
	"voo.su/pkg/gormutil"
)

type MessageRepository struct {
	gormutil.Repo[model.Message]
}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{Repo: gormutil.NewRepo[model.Message](db)}
}
