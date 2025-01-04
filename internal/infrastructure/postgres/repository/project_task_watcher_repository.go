package repository

import (
	"context"
	"gorm.io/gorm"
	"voo.su/internal/infrastructure/postgres/model"
	"voo.su/pkg/gormutil"
)

type ProjectTaskWatcherRepository struct {
	gormutil.Repo[model.ProjectTaskWatcher]
}

func NewProjectTaskWatcherRepository(db *gorm.DB) *ProjectTaskWatcherRepository {
	return &ProjectTaskWatcherRepository{Repo: gormutil.NewRepo[model.ProjectTaskWatcher](db)}
}

func (p *ProjectTaskWatcherRepository) GetWatcherIds(ctx context.Context, taskId int64) []int {
	var ids []int
	_ = p.Repo.Model(ctx).
		Select("member_id").
		Where("task_id = ?", taskId).
		Scan(&ids)

	return ids
}
