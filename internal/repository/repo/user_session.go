package repo

import (
	"gorm.io/gorm"
	"voo.su/internal/repository/model"
	"voo.su/pkg/repo"
)

type UserSession struct {
	repo.Repo[model.UserSession]
}

func NewUserSession(db *gorm.DB) *UserSession {
	return &UserSession{Repo: repo.NewRepo[model.UserSession](db)}
}
