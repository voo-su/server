package repository

import (
	"gorm.io/gorm"
	"voo.su/internal/infrastructure/postgres/model"
	"voo.su/pkg/gormutil"
)

type UserSessionRepository struct {
	gormutil.Repo[model.UserSession]
}

func NewUserSessionRepository(db *gorm.DB) *UserSessionRepository {
	return &UserSessionRepository{Repo: gormutil.NewRepo[model.UserSession](db)}
}
