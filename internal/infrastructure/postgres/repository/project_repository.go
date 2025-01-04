package repository

import (
	"gorm.io/gorm"
	"voo.su/internal/infrastructure/postgres/model"
	"voo.su/pkg/gormutil"
)

type ProjectRepository struct {
	gormutil.Repo[model.Project]
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{Repo: gormutil.NewRepo[model.Project](db)}
}
