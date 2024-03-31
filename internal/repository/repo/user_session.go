package repo

import (
	"gorm.io/gorm"
	"voo.su/internal/repository/model"
	"voo.su/pkg/core"
)

type UserSession struct {
	core.Repo[model.UserSession]
}

func NewUserSession(db *gorm.DB) *UserSession {
	return &UserSession{Repo: core.NewRepo[model.UserSession](db)}
}
