package service

import (
	"context"
	"time"
	"voo.su/internal/repository/model"
)

type ProjectTaskOpt struct {
	ProjectId   int64
	TaskType    int
	Title       string
	Description string
	CreatedBy   int
}

func (p *ProjectService) CreateTask(ctx context.Context, opt *ProjectTaskOpt) (int64, error) {
	task := &model.ProjectTask{
		ProjectId:   opt.ProjectId,
		TaskType:    opt.TaskType,
		Title:       opt.Title,
		Description: opt.Description,
		CreatedBy:   opt.CreatedBy,
		CreatedAt:   time.Now(),
	}

	if err := p.ProjectTask.Create(ctx, task); err != nil {
		return int64(task.Id), err
	}

	return int64(task.Id), nil
}

func (p *ProjectService) TypeTasks(ctx context.Context, projectId int64) ([]*model.ProjectTask, error) {
	query := p.Db().WithContext(ctx).Table("project_task_types")
	query.Where("project_id = ?", projectId)

	var items []*model.ProjectTask
	if err := query.Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (p *ProjectService) Tasks(ctx context.Context, projectId int64, typeId int) ([]*model.ProjectTask, error) {
	query := p.Db().WithContext(ctx).Table("project_tasks")
	query.Where("project_id = ? AND type_id = ?", projectId, typeId)

	var items []*model.ProjectTask
	if err := query.Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}
