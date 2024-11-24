package repo

import (
	"gorm.io/gorm"
	"voo.su/internal/repository/model"
	"voo.su/pkg/repo"
)

type Project struct {
	repo.Repo[model.Project]
}

func NewProject(db *gorm.DB) *Project {
	return &Project{Repo: repo.NewRepo[model.Project](db)}
}
