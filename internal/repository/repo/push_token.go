package repo

import (
	"gorm.io/gorm"
	"voo.su/internal/repository/model"
	"voo.su/pkg/repo"
)

type PushToken struct {
	repo.Repo[model.PushToken]
}

func NewPushToken(db *gorm.DB) *PushToken {
	return &PushToken{Repo: repo.NewRepo[model.PushToken](db)}
}
