package repo

import (
	"gorm.io/gorm"
	"voo.su/internal/repository/model"
	"voo.su/pkg/core"
)

type ProjectTaskComment struct {
	core.Repo[model.ProjectTaskComment]
}

func NewProjectTaskComment(db *gorm.DB) *ProjectTaskComment {
	return &ProjectTaskComment{Repo: core.NewRepo[model.ProjectTaskComment](db)}
}
