package repository

import (
	"gorm.io/gorm"
	"voo.su/internal/infrastructure/postgres/model"
	"voo.su/pkg/gormutil"
)

type MessageLoginRepository struct {
	gormutil.Repo[model.MessageLogin]
}

func NewMessageLoginRepository(db *gorm.DB) *MessageLoginRepository {
	return &MessageLoginRepository{Repo: gormutil.NewRepo[model.MessageLogin](db)}
}
