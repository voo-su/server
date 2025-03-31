package repository

import (
	"gorm.io/gorm"
	"voo.su/internal/infrastructure/postgres/model"
	"voo.su/pkg/gormutil"
)

type MessageForwardedRepository struct {
	gormutil.Repo[model.MessageForwarded]
}

func NewMessageForwardedRepository(db *gorm.DB) *MessageForwardedRepository {
	return &MessageForwardedRepository{Repo: gormutil.NewRepo[model.MessageForwarded](db)}
}
