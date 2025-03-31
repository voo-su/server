package repository

import (
	"gorm.io/gorm"
	"voo.su/internal/infrastructure/postgres/model"
	"voo.su/pkg/gormutil"
)

type MessageLocationRepository struct {
	gormutil.Repo[model.MessageLocation]
}

func NewMessageLocationRepository(db *gorm.DB) *MessageLocationRepository {
	return &MessageLocationRepository{Repo: gormutil.NewRepo[model.MessageLocation](db)}
}
