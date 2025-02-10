package repository

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"voo.su/internal/infrastructure/postgres/model"
	"voo.su/pkg/gormutil"
)

type ProjectTaskCoexecutorRepository struct {
	gormutil.Repo[model.ProjectTaskCoexecutor]
}

func NewProjectTaskCoexecutorRepository(db *gorm.DB) *ProjectTaskCoexecutorRepository {
	return &ProjectTaskCoexecutorRepository{Repo: gormutil.NewRepo[model.ProjectTaskCoexecutor](db)}
}

func (p *ProjectTaskCoexecutorRepository) GetCoexecutorIds(ctx context.Context, taskId uuid.UUID) []int {
	var ids []int
	_ = p.Repo.Model(ctx).
		Select("member_id").
		Where("task_id = ?", taskId).
		Scan(&ids)

	return ids
}
