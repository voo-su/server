package repo

import (
	"gorm.io/gorm"
	"voo.su/internal/repository/model"
	"voo.su/pkg/core"
)

type ProjectTaskType struct {
	core.Repo[model.ProjectTaskType]
}

func NewProjectTaskType(db *gorm.DB) *ProjectTaskType {
	return &ProjectTaskType{Repo: core.NewRepo[model.ProjectTaskType](db)}
}
