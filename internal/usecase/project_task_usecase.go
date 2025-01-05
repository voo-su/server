package usecase

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"time"
	"voo.su/internal/domain/entity"
	postgresModel "voo.su/internal/infrastructure/postgres/model"
)

type ProjectTaskOpt struct {
	ProjectId   int64
	TypeId      int
	Title       string
	Description string
	CreatedBy   int
}

func (p *ProjectUseCase) CreateTask(ctx context.Context, opt *ProjectTaskOpt) (int64, error) {
	task := &postgresModel.ProjectTask{
		ProjectId:   opt.ProjectId,
		TypeId:      opt.TypeId,
		Title:       opt.Title,
		Description: opt.Description,
		AssignerId:  opt.CreatedBy,
		ExecutorId:  opt.CreatedBy,
		CreatedBy:   opt.CreatedBy,
		CreatedAt:   time.Now(),
	}

	if err := p.ProjectTaskRepo.Create(ctx, task); err != nil {
		return task.Id, err
	}

	return task.Id, nil
}

func (p *ProjectUseCase) TypeTasks(ctx context.Context, projectId int64) ([]*postgresModel.ProjectTaskType, error) {
	tx := p.Source.Postgres().WithContext(ctx).
		Table("project_task_types").
		Where("project_id = ?", projectId)

	var items []*postgresModel.ProjectTaskType
	if err := tx.Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (p *ProjectUseCase) Tasks(ctx context.Context, projectId int64, typeId int64) ([]*postgresModel.ProjectTask, error) {
	tx := p.Source.Postgres().WithContext(ctx).
		Table("project_tasks").
		Where("project_id = ? AND type_id = ?", projectId, typeId)

	var items []*postgresModel.ProjectTask
	if err := tx.Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (p *ProjectUseCase) TaskDetail(ctx context.Context, taskId int64) (*entity.ProjectTaskDetailWithMember, error) {
	fields := []string{
		"project_tasks.*",
		"assigner.id AS assigner_id",
		"assigner_user.username AS assigner_username",
		"assigner_user.name AS assigner_name",
		"assigner_user.surname AS assigner_surname",
		"executor_member.id AS executor_member_id",
		"executor_user.username AS executor_username",
		"executor_user.name AS executor_name",
		"executor_user.surname AS executor_surname",
	}

	tx := p.Source.Postgres().WithContext(ctx).
		Table("project_tasks").
		Joins("LEFT JOIN project_members AS assigner ON assigner.id = project_tasks.assigner_id").
		Joins("LEFT JOIN users AS assigner_user ON assigner_user.id = assigner.user_id").
		Joins("LEFT JOIN project_members AS executor_member ON executor_member.id = project_tasks.executor_id").
		Joins("LEFT JOIN users AS executor_user ON executor_user.id = executor_member.user_id").
		Where("project_tasks.id = ?", taskId)

	var taskDetail entity.ProjectTaskDetailWithMember
	if err := tx.Select(fields).Scan(&taskDetail).Error; err != nil {
		return nil, err
	}

	return &taskDetail, nil
}

func (p *ProjectUseCase) TaskExecutor(ctx context.Context, taskId int64, memberId int64) error {
	_, err := p.ProjectTaskRepo.UpdateById(ctx, taskId, map[string]any{
		"executor_id": memberId,
	})
	if err != nil {
		return err
	}

	return nil
}

func (p *ProjectUseCase) TaskMove(ctx context.Context, projectId int64, taskId int64, fromId int64, toId int64) error {
	_, err := p.ProjectTaskRepo.UpdateWhere(ctx, map[string]any{
		"type_id": toId,
	}, "id = ? AND project_id = ? AND type_id = ?", taskId, projectId, fromId)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProjectUseCase) TaskTypeName(ctx context.Context, taskId int64, name string) error {
	_, err := p.ProjectTaskTypeRepo.UpdateWhere(ctx, map[string]any{
		"title": name,
	}, "id = ?", taskId)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProjectUseCase) IsMemberProjectByTask(ctx context.Context, taskId int64, uid int) bool {
	tx := p.Source.Postgres().WithContext(ctx).
		Table("project_tasks").
		Joins("INNER JOIN project_members ON project_members.project_id = project_tasks.project_id").
		Where("project_tasks.id = ? AND project_members.user_id = ?", taskId, uid)

	var count int64
	if err := tx.Count(&count).Error; err != nil {
		return false
	}

	return count > 0
}

func (p *ProjectUseCase) InviteCoexecutor(ctx context.Context, taskId int64, memberIds []int, uid int) error {
	var (
		err            error
		addCoexecutors []*postgresModel.ProjectTaskCoexecutor
		db             = p.Source.Postgres().WithContext(ctx)
	)
	m := make(map[int]struct{})
	for _, value := range p.ProjectTaskCoexecutorRepo.GetCoexecutorIds(ctx, taskId) {
		m[value] = struct{}{}
	}

	for _, value := range memberIds {
		if _, ok := m[value]; !ok {
			addCoexecutors = append(addCoexecutors, &postgresModel.ProjectTaskCoexecutor{
				TaskId:    int(taskId),
				MemberId:  value,
				CreatedBy: uid,
			})
		}
	}

	if len(addCoexecutors) == 0 {
		return errors.New(p.Locale.Localize("all_invited_contacts_are_task_coexecutors"))
	}

	if err = db.Transaction(func(tx *gorm.DB) error {
		if err = tx.Create(&addCoexecutors).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (p *ProjectUseCase) GetCoexecutors(ctx context.Context, taskId int64) ([]*entity.ProjectMemberItem, error) {
	fields := []string{
		"project_members.id AS id",
		"project_members.user_id AS user_id",
		"users.username AS username",
		"users.name AS Name",
		"users.surname AS Surname",
	}
	tx := p.Source.Postgres().WithContext(ctx).
		Table("project_task_coexecutors").
		Joins("LEFT JOIN project_members ON project_members.id = project_task_coexecutors.member_id").
		Joins("LEFT JOIN users ON users.id = project_task_coexecutors.member_id").
		Where("project_task_coexecutors.task_id = ?", taskId)

	var items []*entity.ProjectMemberItem
	if err := tx.Unscoped().Select(fields).Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (p *ProjectUseCase) InviteWatcher(ctx context.Context, taskId int64, memberIds []int, uid int) error {
	var (
		err         error
		addWatchers []*postgresModel.ProjectTaskWatcher
		db          = p.Source.Postgres().WithContext(ctx)
	)
	m := make(map[int]struct{})
	for _, value := range p.ProjectTaskWatcherRepo.GetWatcherIds(ctx, taskId) {
		m[value] = struct{}{}
	}

	for _, value := range memberIds {
		if _, ok := m[value]; !ok {
			addWatchers = append(addWatchers, &postgresModel.ProjectTaskWatcher{
				TaskId:    int(taskId),
				MemberId:  value,
				CreatedBy: uid,
			})
		}
	}

	if len(addWatchers) == 0 {
		return errors.New(p.Locale.Localize("all_invited_contacts_are_task_observers"))
	}

	if err = db.Transaction(func(tx *gorm.DB) error {
		if err = tx.Create(&addWatchers).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (p *ProjectUseCase) GetWatchers(ctx context.Context, taskId int64) ([]*entity.ProjectMemberItem, error) {
	fields := []string{
		"project_members.id AS id",
		"project_members.user_id AS user_id",
		"users.username AS username",
		"users.name AS Name",
		"users.surname AS Surname",
	}
	tx := p.Source.Postgres().WithContext(ctx).
		Table("project_task_watchers").
		Joins("LEFT JOIN project_members ON project_members.id = project_task_watchers.member_id").
		Joins("LEFT JOIN users ON users.id = project_task_watchers.member_id").
		Where("project_task_watchers.task_id = ?", taskId)

	var items []*entity.ProjectMemberItem
	if err := tx.Unscoped().Select(fields).Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

type ProjectCommentOpt struct {
	TaskId    int64
	Comment   string
	CreatedBy int
}

func (p *ProjectUseCase) CreateComment(ctx context.Context, opt *ProjectCommentOpt) (int64, error) {
	comment := &postgresModel.ProjectTaskComment{
		TaskId:    opt.TaskId,
		Comment:   opt.Comment,
		CreatedBy: opt.CreatedBy,
		CreatedAt: time.Now(),
	}

	err := p.ProjectTaskCommentRepo.Create(ctx, comment)
	if err != nil {
		return comment.Id, err
	}

	return comment.Id, nil
}

func (p *ProjectUseCase) Comments(ctx context.Context, TaskId int64) ([]*postgresModel.ProjectTaskComment, error) {
	tx := p.Source.Postgres().WithContext(ctx).
		Table("project_task_comments").
		Where("task_id = ?", TaskId)

	var items []*postgresModel.ProjectTaskComment
	if err := tx.Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}
