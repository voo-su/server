package service

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"time"
	"voo.su/internal/repository/cache"
	"voo.su/internal/repository/model"
	"voo.su/internal/repository/repo"
)

type ProjectService struct {
	*repo.Source
	ProjectRepo            *repo.Project
	ProjectMemberRepo      *repo.ProjectMember
	ProjectTaskTypeRepo    *repo.ProjectTaskType
	ProjectTaskRepo        *repo.ProjectTask
	ProjectTaskCommentRepo *repo.ProjectTaskComment
	UserRepo               *repo.User
	Relation               *cache.Relation
}

func NewProjectService(
	source *repo.Source,
	projectRepo *repo.Project,
	projectMemberRepo *repo.ProjectMember,
	projectTaskTypeRepo *repo.ProjectTaskType,
	projectTaskRepo *repo.ProjectTask,
	projectTaskCommentRepo *repo.ProjectTaskComment,
	userRepo *repo.User,
	relation *cache.Relation,
) *ProjectService {
	return &ProjectService{
		Source:                 source,
		ProjectRepo:            projectRepo,
		ProjectMemberRepo:      projectMemberRepo,
		ProjectTaskTypeRepo:    projectTaskTypeRepo,
		ProjectTaskRepo:        projectTaskRepo,
		ProjectTaskCommentRepo: projectTaskCommentRepo,
		UserRepo:               userRepo,
		Relation:               relation,
	}
}

type ProjectOpt struct {
	UserId int
	Title  string
}

func (p *ProjectService) CreateProject(ctx context.Context, opt *ProjectOpt) (int64, error) {
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

func (p *ProjectService) Projects(ctx context.Context) ([]*model.Project, error) {
	query := p.Db().WithContext(ctx).Table("projects")

	var items []*model.Project
	if err := query.Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (p *ProjectService) GetMembers(ctx context.Context, projectId int64) []*model.ProjectMemberItem {
	fields := []string{
		"project_members.id",
		"project_members.user_id",
		"users.username",
	}
	tx := p.Db().WithContext(ctx).Table("project_members")
	tx.Joins("LEFT JOIN users on users.id = project_members.user_id")
	tx.Where("project_members.project_id = ?", projectId)
	tx.Order("project_members.leader desc")
	var items []*model.ProjectMemberItem
	tx.Unscoped().Select(fields).Scan(&items)

	return items
}

func (p *ProjectService) IsMember(ctx context.Context, gid, uid int, cache bool) bool {
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

func (p *ProjectService) Invite(ctx context.Context, opt *ProjectInviteOpt) error {
	var (
		err        error
		addMembers []*model.ProjectMember
		db         = p.Source.Db().WithContext(ctx)
	)
	m := make(map[int]struct{})
	for _, value := range p.ProjectMemberRepo.GetMemberIds(ctx, opt.ProjectId) {
		m[value] = struct{}{}
	}

	mids := make([]int, 0)
	mids = append(mids, opt.MemberIds...)
	mids = append(mids, opt.UserId)

	memberItems := make([]*model.User, 0)
	err = db.Table("users").
		Select("id, username").
		Where("id in ?", mids).
		Scan(&memberItems).Error
	if err != nil {
		return err
	}

	memberMaps := make(map[int]*model.User)
	for _, item := range memberItems {
		memberMaps[item.Id] = item
	}

	members := make([]model.ProjectMemberItem, 0)
	for _, value := range opt.MemberIds {
		members = append(members, model.ProjectMemberItem{
			UserId:   int64(value),
			Username: memberMaps[value].Username,
		})
		if _, ok := m[value]; !ok {
			addMembers = append(addMembers, &model.ProjectMember{
				ProjectId: opt.ProjectId,
				UserId:    value,
				CreatedBy: opt.UserId,
			})
		}
	}
	if len(addMembers) == 0 {
		return errors.New("все приглашенные контакты стали участниками проекта")
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		tx.Delete(&model.ProjectMember{}, "project_id = ? and user_id in ?", opt.ProjectId, opt.MemberIds)
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
