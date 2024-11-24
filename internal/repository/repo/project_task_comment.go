package repo

import (
	"gorm.io/gorm"
	"voo.su/internal/repository/model"
	"voo.su/pkg/repo"
)

type ProjectTaskComment struct {
	repo.Repo[model.ProjectTaskComment]
}

func NewProjectTaskComment(db *gorm.DB) *ProjectTaskComment {
	return &ProjectTaskComment{Repo: repo.NewRepo[model.ProjectTaskComment](db)}
}
