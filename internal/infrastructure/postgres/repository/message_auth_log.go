package repository

import (
	"gorm.io/gorm"
	"voo.su/internal/infrastructure/postgres/model"
	"voo.su/pkg/gormutil"
)

type MessageAuthLogRepository struct {
	gormutil.Repo[model.MessageAuthLog]
}

func NewMessageAuthLogForwardedRepository(db *gorm.DB) *MessageAuthLogRepository {
	return &MessageAuthLogRepository{Repo: gormutil.NewRepo[model.MessageAuthLog](db)}
}
