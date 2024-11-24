package repo

import (
	"gorm.io/gorm"
	"voo.su/internal/repository/model"
	"voo.su/pkg/repo"
)

type ProjectTaskType struct {
	repo.Repo[model.ProjectTaskType]
}

func NewProjectTaskType(db *gorm.DB) *ProjectTaskType {
	return &ProjectTaskType{Repo: repo.NewRepo[model.ProjectTaskType](db)}
}
