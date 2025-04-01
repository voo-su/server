package repository

import (
	"gorm.io/gorm"
	"voo.su/internal/infrastructure/postgres/model"
	"voo.su/pkg/gormutil"
)

type MessageSystemRepository struct {
	gormutil.Repo[model.MessageSystem]
}

func NewMessageSystemRepository(db *gorm.DB) *MessageSystemRepository {
	return &MessageSystemRepository{Repo: gormutil.NewRepo[model.MessageSystem](db)}
}
