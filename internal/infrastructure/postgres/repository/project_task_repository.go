// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package repository

import (
	"gorm.io/gorm"
	"voo.su/internal/infrastructure/postgres/model"
	"voo.su/pkg/repo"
)

type ProjectTaskRepository struct {
	repo.Repo[model.ProjectTask]
}

func NewProjectTaskRepository(db *gorm.DB) *ProjectTaskRepository {
	return &ProjectTaskRepository{Repo: repo.NewRepo[model.ProjectTask](db)}
}
