package service

import (
	"context"
	"time"
	"voo.su/internal/repository/model"
	"voo.su/internal/repository/repo"
)

type ProjectService struct {
	*repo.Source
	ProjectRepo            *repo.Project
	ProjectTaskTypeRepo    *repo.ProjectTaskType
	ProjectTaskRepo        *repo.ProjectTask
	ProjectTaskCommentRepo *repo.ProjectTaskComment
	UserRepo               *repo.User
}

func NewProjectService(
	source *repo.Source,
	projectRepo *repo.Project,
	projectTaskTypeRepo *repo.ProjectTaskType,
	projectTaskRepo *repo.ProjectTask,
	projectTaskCommentRepo *repo.ProjectTaskComment,
	userRepo *repo.User,
) *ProjectService {
	return &ProjectService{
		Source:                 source,
		ProjectRepo:            projectRepo,
		ProjectTaskTypeRepo:    projectTaskTypeRepo,
		ProjectTaskRepo:        projectTaskRepo,
		ProjectTaskCommentRepo: projectTaskCommentRepo,
		UserRepo:               userRepo,
	}
}

type ProjectOpt struct {
	UserId int
	Title  string
}

func (p *ProjectService) CreateProject(ctx context.Context, opt *ProjectOpt) (int64, error) {
	project := &model.Project{
		Name:      opt.Title,
		CreatedBy: opt.UserId,
		CreatedAt: time.Now(),
	}

	if err := p.ProjectRepo.Create(ctx, project); err != nil {
		return int64(project.Id), err
	}

	var projectTaskTypes = &[]model.ProjectTaskType{
		{ProjectId: project.Id, Title: "Новые", CreatedBy: opt.UserId},
		{ProjectId: project.Id, Title: "Выполняются", CreatedBy: opt.UserId},
		{ProjectId: project.Id, Title: "Тест", CreatedBy: opt.UserId},
		{ProjectId: project.Id, Title: "Сделаны", CreatedBy: opt.UserId},
	}

	err := p.ProjectTaskTypeRepo.Db.WithContext(ctx).Create(&projectTaskTypes).Error
	if err != nil {
		return int64(project.Id), err
	}

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
