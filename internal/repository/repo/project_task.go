package repo

import (
	"gorm.io/gorm"
	"voo.su/internal/repository/model"
	"voo.su/pkg/core"
)

type ProjectTask struct {
	core.Repo[model.ProjectTask]
}

func NewProjectTask(db *gorm.DB) *ProjectTask {
	return &ProjectTask{Repo: core.NewRepo[model.ProjectTask](db)}
}
