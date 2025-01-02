// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

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
