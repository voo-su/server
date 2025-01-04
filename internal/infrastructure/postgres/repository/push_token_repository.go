package repository

import (
	"gorm.io/gorm"
	"voo.su/internal/infrastructure/postgres/model"
	"voo.su/pkg/gormutil"
)

type PushTokenRepository struct {
	gormutil.Repo[model.PushToken]
}

func NewPushTokenRepository(db *gorm.DB) *PushTokenRepository {
	return &PushTokenRepository{Repo: gormutil.NewRepo[model.PushToken](db)}
}
