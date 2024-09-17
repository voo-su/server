package repo

import (
	"gorm.io/gorm"
	"voo.su/internal/repository/model"
	"voo.su/pkg/core"
)

type ProjectTaskWatcher struct {
	core.Repo[model.ProjectTaskWatcher]
}

func NewProjectTaskWatcher(db *gorm.DB) *ProjectTaskWatcher {
	return &ProjectTaskWatcher{Repo: core.NewRepo[model.ProjectTaskWatcher](db)}
}
