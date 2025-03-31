package repository

import (
	"gorm.io/gorm"
	"voo.su/internal/infrastructure/postgres/model"
	"voo.su/pkg/gormutil"
)

type MessageMediaRepository struct {
	gormutil.Repo[model.MessageMedia]
}

func NewMessageMediaRepository(db *gorm.DB) *MessageMediaRepository {
	return &MessageMediaRepository{Repo: gormutil.NewRepo[model.MessageMedia](db)}
}
