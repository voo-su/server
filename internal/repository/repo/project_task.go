package repo

import (
	"gorm.io/gorm"
	"voo.su/internal/repository/model"
	"voo.su/pkg/repo"
)

type ProjectTask struct {
	repo.Repo[model.ProjectTask]
}

func NewProjectTask(db *gorm.DB) *ProjectTask {
	return &ProjectTask{Repo: repo.NewRepo[model.ProjectTask](db)}
}
