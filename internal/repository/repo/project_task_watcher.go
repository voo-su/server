package repo

import (
	"context"
	"gorm.io/gorm"
	"voo.su/internal/repository/model"
	"voo.su/pkg/core"
)

type ProjectTaskWatcher struct {
	core.Repo[model.ProjectTaskWatcher]
}

func NewProjectTaskWatcher(db *gorm.DB) *ProjectTaskWatcher {
	return &ProjectTaskWatcher{Repo: core.NewRepo[model.ProjectTaskWatcher](db)}
}

func (p *ProjectTaskWatcher) GetWatcherIds(ctx context.Context, taskId int64) []int {
	var ids []int
	_ = p.Repo.Model(ctx).Select("member_id").
		Where("task_id = ?", taskId).
		Scan(&ids)

	return ids
}
