package repository

import (
	"gorm.io/gorm"
	"voo.su/internal/infrastructure/postgres/model"
	"voo.su/pkg/repo"
)

type UserSessionRepository struct {
	repo.Repo[model.UserSession]
}

func NewUserSessionRepository(db *gorm.DB) *UserSessionRepository {
	return &UserSessionRepository{Repo: repo.NewRepo[model.UserSession](db)}
}
