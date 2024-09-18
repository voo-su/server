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

	if err := p.ProjectTaskRepo.Create(ctx, task); err != nil {
		return task.Id, err
	}

	return task.Id, nil
}

func (p *ProjectService) TypeTasks(ctx context.Context, projectId int64) ([]*model.ProjectTaskType, error) {
	tx := p.Db().WithContext(ctx).Table("project_task_types")
	tx.Where("project_id = ?", projectId)

	var items []*model.ProjectTaskType
	if err := tx.Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (p *ProjectService) Tasks(ctx context.Context, projectId int64, typeId int64) ([]*model.ProjectTask, error) {
	tx := p.Db().WithContext(ctx).Table("project_tasks")
	tx.Where("project_id = ? AND type_id = ?", projectId, typeId)

	var items []*model.ProjectTask
	if err := tx.Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (p *ProjectService) TaskDetail(ctx context.Context, taskId int64) (*model.ProjectTaskDetailWithMember, error) {
	fields := []string{
		"project_tasks.*",
		"assigner.id AS assigner_id",
		"assigner_user.username AS assigner_username",
		"executor_member.id AS executor_member_id",
		"executor_user.username AS executor_username",
	}

	tx := p.Db().WithContext(ctx).Table("project_tasks").
		Joins("LEFT JOIN project_members AS assigner ON assigner.id = project_tasks.assigner_id").
		Joins("LEFT JOIN users AS assigner_user ON assigner_user.id = assigner.user_id").
		Joins("LEFT JOIN project_members AS executor_member ON executor_member.id = project_tasks.executor_id").
		Joins("LEFT JOIN users AS executor_user ON executor_user.id = executor_member.user_id").
		Where("project_tasks.id = ?", taskId)

	var taskDetail model.ProjectTaskDetailWithMember
	if err := tx.Select(fields).Scan(&taskDetail).Error; err != nil {
		return nil, err
	}

	return &taskDetail, nil
}

func (p *ProjectService) TaskMove(ctx context.Context, projectId int64, taskId int64, fromId int64, toId int64) error {
	_, err := p.ProjectTaskRepo.UpdateWhere(ctx, map[string]any{
		"type_id": toId,
	}, "id = ? AND project_id = ? AND type_id = ?", taskId, projectId, fromId)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProjectService) TaskTypeName(ctx context.Context, taskId int64, name string) error {
	_, err := p.ProjectTaskTypeRepo.UpdateWhere(ctx, map[string]any{
		"title": name,
	}, "id = ?", taskId)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProjectService) GetCoexecutors(ctx context.Context, taskId int64) ([]*model.ProjectMemberItem, error) {
	fields := []string{
		"project_members.id AS id",
		"project_members.user_id AS user_id",
		"users.username AS username",
	}
	tx := p.Db().WithContext(ctx).Table("project_task_coexecutors").
		Joins("LEFT JOIN project_members ON project_members.id = project_task_coexecutors.member_id").
		Joins("LEFT JOIN users ON users.id = project_task_coexecutors.member_id").
		Where("project_task_coexecutors.task_id = ?", taskId)

	var items []*model.ProjectMemberItem
	if err := tx.Unscoped().Select(fields).Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (p *ProjectService) GetWatchers(ctx context.Context, taskId int64) ([]*model.ProjectMemberItem, error) {
	fields := []string{
		"project_members.id AS id",
		"project_members.user_id AS user_id",
		"users.username AS username",
	}
	tx := p.Db().WithContext(ctx).Table("project_task_watchers").
		Joins("LEFT JOIN project_members ON project_members.id = project_task_watchers.member_id").
		Joins("LEFT JOIN users ON users.id = project_task_watchers.member_id").
		Where("project_task_watchers.task_id = ?", taskId)

	var items []*model.ProjectMemberItem
	if err := tx.Unscoped().Select(fields).Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}
