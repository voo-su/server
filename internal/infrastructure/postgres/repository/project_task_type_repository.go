package repository

import (
	"gorm.io/gorm"
	"voo.su/internal/infrastructure/postgres/model"
	"voo.su/pkg/gormutil"
)

type ProjectTaskTypeRepository struct {
	gormutil.Repo[model.ProjectTaskType]
}

func NewProjectTaskTypeRepository(db *gorm.DB) *ProjectTaskTypeRepository {
	return &ProjectTaskTypeRepository{Repo: gormutil.NewRepo[model.ProjectTaskType](db)}
}
