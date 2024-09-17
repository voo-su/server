package repo

import (
	"gorm.io/gorm"
	"voo.su/internal/repository/model"
	"voo.su/pkg/core"
)

type ProjectTaskCoexecutor struct {
	core.Repo[model.ProjectTaskCoexecutor]
}

func NewProjectTaskCoexecutor(db *gorm.DB) *ProjectTaskCoexecutor {
	return &ProjectTaskCoexecutor{Repo: core.NewRepo[model.ProjectTaskCoexecutor](db)}
}
