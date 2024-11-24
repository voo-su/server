package usecase

import (
	"context"
	"errors"
	"gorm.io/gorm"
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

func (p *ProjectUseCase) CreateTask(ctx context.Context, opt *ProjectTaskOpt) (int64, error) {
	task := &model.ProjectTask{
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

func (p *ProjectUseCase) TypeTasks(ctx context.Context, projectId int64) ([]*model.ProjectTaskType, error) {
	tx := p.Db().WithContext(ctx).Table("project_task_types")
	tx.Where("project_id = ?", projectId)

	var items []*model.ProjectTaskType
	if err := tx.Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (p *ProjectUseCase) Tasks(ctx context.Context, projectId int64, typeId int64) ([]*model.ProjectTask, error) {
	tx := p.Db().WithContext(ctx).Table("project_tasks")
	tx.Where("project_id = ? AND type_id = ?", projectId, typeId)

	var items []*model.ProjectTask
	if err := tx.Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (p *ProjectUseCase) TaskDetail(ctx context.Context, taskId int64) (*model.ProjectTaskDetailWithMember, error) {
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
	tx := p.Db().WithContext(ctx).
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
		addCoexecutors []*model.ProjectTaskCoexecutor
		db             = p.Source.Db().WithContext(ctx)
	)
	m := make(map[int]struct{})
	for _, value := range p.ProjectTaskCoexecutorRepo.GetCoexecutorIds(ctx, taskId) {
		m[value] = struct{}{}
	}

	for _, value := range memberIds {
		if _, ok := m[value]; !ok {
			addCoexecutors = append(addCoexecutors, &model.ProjectTaskCoexecutor{
				TaskId:    int(taskId),
				MemberId:  value,
				CreatedBy: uid,
			})
		}
	}

	if len(addCoexecutors) == 0 {
		return errors.New("все приглашённые контакты стали соисполнителями задачи")
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		if err = tx.Create(&addCoexecutors).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (p *ProjectUseCase) GetCoexecutors(ctx context.Context, taskId int64) ([]*model.ProjectMemberItem, error) {
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

func (p *ProjectUseCase) InviteWatcher(ctx context.Context, taskId int64, memberIds []int, uid int) error {
	var (
		err         error
		addWatchers []*model.ProjectTaskWatcher
		db          = p.Source.Db().WithContext(ctx)
	)
	m := make(map[int]struct{})
	for _, value := range p.ProjectTaskWatcherRepo.GetWatcherIds(ctx, taskId) {
		m[value] = struct{}{}
	}

	for _, value := range memberIds {
		if _, ok := m[value]; !ok {
			addWatchers = append(addWatchers, &model.ProjectTaskWatcher{
				TaskId:    int(taskId),
				MemberId:  value,
				CreatedBy: uid,
			})
		}
	}

	if len(addWatchers) == 0 {
		return errors.New("все приглашённые контакты стали наблюдателями задачи")
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		if err = tx.Create(&addWatchers).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (p *ProjectUseCase) GetWatchers(ctx context.Context, taskId int64) ([]*model.ProjectMemberItem, error) {
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
