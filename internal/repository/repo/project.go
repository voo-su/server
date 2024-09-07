package repo

import (
	"gorm.io/gorm"
	"voo.su/internal/repository/model"
	"voo.su/pkg/core"
)

type Project struct {
	core.Repo[model.Project]
}

func NewProject(db *gorm.DB) *Project {
	return &Project{Repo: core.NewRepo[model.Project](db)}
}
