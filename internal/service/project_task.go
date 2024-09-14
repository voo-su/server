package service

import (
	"context"
	"time"
	"voo.su/internal/repository/model"
)

type ProjectTaskOpt struct {
	ProjectId   int64
	TypeId      int
	Title       string
	Description string
	CreatedBy   int
}

func (p *ProjectService) CreateTask(ctx context.Context, opt *ProjectTaskOpt) (int64, error) {
	task := &model.ProjectTask{
		ProjectId:   opt.ProjectId,
		TypeId:      opt.TypeId,
		Title:       opt.Title,
		Description: opt.Description,
		CreatedBy:   opt.CreatedBy,
		CreatedAt:   time.Now(),
	}

	err := p.ProjectTask.Create(ctx, task)
	if err != nil {
		return task.Id, err
	}

	return task.Id, nil
}

func (p *ProjectService) TypeTasks(ctx context.Context, projectId int64) ([]*model.ProjectTaskType, error) {
	query := p.Db().WithContext(ctx).Table("project_task_types")
	query.Where("project_id = ?", projectId)

	var items []*model.ProjectTaskType
	if err := query.Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (p *ProjectService) Tasks(ctx context.Context, projectId int64, typeId int64) ([]*model.ProjectTask, error) {
	query := p.Db().WithContext(ctx).Table("project_tasks")
	query.Where("project_id = ? AND type_id = ?", projectId, typeId)

	var items []*model.ProjectTask
	if err := query.Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (p *ProjectService) TaskMove(ctx context.Context, projectId int64, taskId int64, fromId int64, toId int64) error {
	_, err := p.ProjectTask.UpdateWhere(ctx, map[string]any{
		"type_id": toId,
	}, "id = ? AND project_id = ? AND type_id = ?", taskId, projectId, fromId)

	if err != nil {
		return err
	}

	return nil
}
