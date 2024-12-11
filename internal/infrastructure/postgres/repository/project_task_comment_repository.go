package repository

import (
	"gorm.io/gorm"
	"voo.su/internal/infrastructure/postgres/model"
	"voo.su/pkg/repo"
)

type ProjectTaskCommentRepository struct {
	repo.Repo[model.ProjectTaskComment]
}

func NewProjectTaskCommentRepository(db *gorm.DB) *ProjectTaskCommentRepository {
	return &ProjectTaskCommentRepository{Repo: repo.NewRepo[model.ProjectTaskComment](db)}
}
