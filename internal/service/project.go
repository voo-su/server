package service

import (
	"context"
	"time"
	"voo.su/internal/repository/model"
	"voo.su/internal/repository/repo"
)

type ProjectService struct {
	*repo.Source
	Project            *repo.Project
	ProjectTask        *repo.ProjectTask
	ProjectTaskComment *repo.ProjectTaskComment
}

func NewProjectService(
	source *repo.Source,
	project *repo.Project,
	projectTask *repo.ProjectTask,
	projectTaskComment *repo.ProjectTaskComment,
) *ProjectService {
	return &ProjectService{
		Source:             source,
		Project:            project,
		ProjectTask:        projectTask,
		ProjectTaskComment: projectTaskComment,
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

	err := p.Project.Create(ctx, project)
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
