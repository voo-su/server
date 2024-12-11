package repository

import (
	"gorm.io/gorm"
	"voo.su/internal/infrastructure/postgres/model"
	"voo.su/pkg/repo"
)

type ProjectTaskTypeRepository struct {
	repo.Repo[model.ProjectTaskType]
}

func NewProjectTaskTypeRepository(db *gorm.DB) *ProjectTaskTypeRepository {
	return &ProjectTaskTypeRepository{Repo: repo.NewRepo[model.ProjectTaskType](db)}
}
