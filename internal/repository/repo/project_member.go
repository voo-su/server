package repo

import (
	"gorm.io/gorm"
	"voo.su/internal/repository/model"
	"voo.su/pkg/core"
)

type ProjectMember struct {
	core.Repo[model.ProjectMember]
}

func NewProjectMember(db *gorm.DB) *ProjectMember {
	return &ProjectMember{Repo: core.NewRepo[model.ProjectMember](db)}
}
