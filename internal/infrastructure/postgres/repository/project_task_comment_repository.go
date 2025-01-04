package repository

import (
	"gorm.io/gorm"
	"voo.su/internal/infrastructure/postgres/model"
	"voo.su/pkg/gormutil"
)

type ProjectTaskCommentRepository struct {
	gormutil.Repo[model.ProjectTaskComment]
}

func NewProjectTaskCommentRepository(db *gorm.DB) *ProjectTaskCommentRepository {
	return &ProjectTaskCommentRepository{Repo: gormutil.NewRepo[model.ProjectTaskComment](db)}
}
