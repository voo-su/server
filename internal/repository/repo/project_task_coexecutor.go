package repo

import (
	"context"
	"gorm.io/gorm"
	"voo.su/internal/repository/model"
	"voo.su/pkg/core"
)

type ProjectTaskCoexecutor struct {
	core.Repo[model.ProjectTaskCoexecutor]
}

func NewProjectTaskCoexecutor(db *gorm.DB) *ProjectTaskCoexecutor {
	return &ProjectTaskCoexecutor{Repo: core.NewRepo[model.ProjectTaskCoexecutor](db)}
}

func (p *ProjectTaskCoexecutor) GetCoexecutorIds(ctx context.Context, taskId int64) []int {
	var ids []int
	_ = p.Repo.Model(ctx).Select("member_id").
		Where("task_id = ?", taskId).
		Scan(&ids)

	return ids
}
