// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package repository

import (
	"context"
	"gorm.io/gorm"
	"voo.su/internal/infrastructure/postgres/model"
	"voo.su/pkg/repo"
)

type ProjectTaskCoexecutorRepository struct {
	repo.Repo[model.ProjectTaskCoexecutor]
}

func NewProjectTaskCoexecutorRepository(db *gorm.DB) *ProjectTaskCoexecutorRepository {
	return &ProjectTaskCoexecutorRepository{Repo: repo.NewRepo[model.ProjectTaskCoexecutor](db)}
}

func (p *ProjectTaskCoexecutorRepository) GetCoexecutorIds(ctx context.Context, taskId int64) []int {
	var ids []int
	_ = p.Repo.Model(ctx).
		Select("member_id").
		Where("task_id = ?", taskId).
		Scan(&ids)

	return ids
}
