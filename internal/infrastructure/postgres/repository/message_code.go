package repository

import (
	"gorm.io/gorm"
	"voo.su/internal/infrastructure/postgres/model"
	"voo.su/pkg/gormutil"
)

type MessageCodeRepository struct {
	gormutil.Repo[model.MessageCode]
}

func NewMessageCodeRepository(db *gorm.DB) *MessageCodeRepository {
	return &MessageCodeRepository{Repo: gormutil.NewRepo[model.MessageCode](db)}
}
