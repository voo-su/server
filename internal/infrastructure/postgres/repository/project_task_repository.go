package repository

import (
	"gorm.io/gorm"
	"voo.su/internal/infrastructure/postgres/model"
	"voo.su/pkg/gormutil"
)

type ProjectTaskRepository struct {
	gormutil.Repo[model.ProjectTask]
}

func NewProjectTaskRepository(db *gorm.DB) *ProjectTaskRepository {
	return &ProjectTaskRepository{Repo: gormutil.NewRepo[model.ProjectTask](db)}
}
