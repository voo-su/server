package usecase

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"time"
	"voo.su/internal/repository/cache"
	"voo.su/internal/repository/model"
	"voo.su/internal/repository/repo"
)

type ProjectUseCase struct {
	*repo.Source
	ProjectRepo               *repo.Project
	ProjectMemberRepo         *repo.ProjectMember
	ProjectTaskTypeRepo       *repo.ProjectTaskType
	ProjectTaskRepo           *repo.ProjectTask
	ProjectTaskCommentRepo    *repo.ProjectTaskComment
	UserRepo                  *repo.User
	Relation                  *cache.Relation
	ProjectTaskCoexecutorRepo *repo.ProjectTaskCoexecutor
	ProjectTaskWatcherRepo    *repo.ProjectTaskWatcher
}

func NewProjectUseCase(
	source *repo.Source,
	projectRepo *repo.Project,
	projectMemberRepo *repo.ProjectMember,
	projectTaskTypeRepo *repo.ProjectTaskType,
	projectTaskRepo *repo.ProjectTask,
	projectTaskCommentRepo *repo.ProjectTaskComment,
	userRepo *repo.User,
	relation *cache.Relation,
	projectTaskCoexecutorRepo *repo.ProjectTaskCoexecutor,
	projectTaskWatcherRepo *repo.ProjectTaskWatcher,
) *ProjectUseCase {
	return &ProjectUseCase{
		Source:                    source,
		ProjectRepo:               projectRepo,
		ProjectMemberRepo:         projectMemberRepo,
		ProjectTaskTypeRepo:       projectTaskTypeRepo,
		ProjectTaskRepo:           projectTaskRepo,
		ProjectTaskCommentRepo:    projectTaskCommentRepo,
		UserRepo:                  userRepo,
		Relation:                  relation,
		ProjectTaskCoexecutorRepo: projectTaskCoexecutorRepo,
		ProjectTaskWatcherRepo:    projectTaskWatcherRepo,
	}
}

type ProjectOpt struct {
	UserId int
	Title  string
}

func (p *ProjectUseCase) CreateProject(ctx context.Context, opt *ProjectOpt) (int64, error) {
	var (
		err       error
		members   []*model.ProjectMember
		taskTypes []*model.ProjectTaskType
	)

	project := &model.Project{
		Name:      opt.Title,
		CreatedBy: opt.UserId,
		CreatedAt: time.Now(),
	}

	types := [4]string{"Новые", "Выполняются", "Тест", "Сделаны"}

	err = p.Source.Db().Transaction(func(tx *gorm.DB) error {
		if err = tx.Create(project).Error; err != nil {
			return err
		}

		members = append(members, &model.ProjectMember{
			ProjectId: project.Id,
			UserId:    opt.UserId,
			CreatedBy: opt.UserId,
		})

		if err = tx.Create(members).Error; err != nil {
			return err
		}

		for _, title := range types {
			taskTypes = append(taskTypes, &model.ProjectTaskType{
				ProjectId: project.Id,
				Title:     title,
				CreatedBy: opt.UserId,
			})
		}
		if err = tx.Create(taskTypes).Error; err != nil {
			return err
		}

		return nil
	})

	return int64(project.Id), nil
}

func (p *ProjectUseCase) Projects(userId int) ([]*model.ProjectItem, error) {
	tx := p.Source.Db().Table("project_members")
	tx.Select("p.id AS id, p.name AS name")
	tx.Joins("LEFT JOIN projects p ON p.id = project_members.project_id")
	tx.Where("project_members.user_id = ?", userId)
	tx.Order("project_members.created_at desc")

	items := make([]*model.ProjectItem, 0)
	if err := tx.Scan(&items).Error; err != nil {
		return nil, err
	}

	length := len(items)
	if length == 0 {
		return items, nil
	}

	ids := make([]int, 0, length)
	for i := range items {
		ids = append(ids, items[i].Id)
	}

	return items, nil
}

func (p *ProjectUseCase) GetMembers(ctx context.Context, projectId int64) []*model.ProjectMemberItem {
	fields := []string{
		"project_members.id AS id",
		"project_members.user_id AS user_id",
		"users.username AS username",
	}
	tx := p.Db().WithContext(ctx).Table("project_members").
		Joins("LEFT JOIN users ON users.id = project_members.user_id").
		Where("project_members.project_id = ?", projectId)
	//.Order("project_members.leader desc")

	var items []*model.ProjectMemberItem
	tx.Unscoped().Select(fields).Scan(&items)

	return items
}

func (p *ProjectUseCase) IsMember(ctx context.Context, gid, uid int, cache bool) bool {
	if cache && p.Relation.IsGroupRelation(ctx, uid, gid) == nil {
		return true
	}

	exist, err := p.ProjectMemberRepo.QueryExist(ctx, "project_id = ? and user_id = ?", gid, uid)
	if err != nil {
		return false
	}
	if exist {
		p.Relation.SetGroupRelation(ctx, uid, gid)
	}

	return exist
}

type ProjectInviteOpt struct {
	ProjectId int
	UserId    int
	MemberIds []int
}

func (p *ProjectUseCase) Invite(ctx context.Context, opt *ProjectInviteOpt) error {
	var (
		err        error
		addMembers []*model.ProjectMember
		db         = p.Source.Db().WithContext(ctx)
	)
	m := make(map[int]struct{})
	for _, value := range p.ProjectMemberRepo.GetMemberIds(ctx, opt.ProjectId) {
		m[value] = struct{}{}
	}

	for _, value := range opt.MemberIds {
		if _, ok := m[value]; !ok {
			addMembers = append(addMembers, &model.ProjectMember{
				ProjectId: opt.ProjectId,
				UserId:    value,
				CreatedBy: opt.UserId,
			})
		}
	}
	if len(addMembers) == 0 {
		return errors.New("все приглашённые контакты уже являются участниками проекта")
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		tx.Delete(&model.ProjectMember{}, "project_id = ? AND user_id in ?", opt.ProjectId, opt.MemberIds)
		if err = tx.Create(&addMembers).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
