package service

import (
	"context"
	"time"
	"voo.su/internal/repository/model"
)

type ProjectCommentOpt struct {
	TaskId    int64
	Comment   string
	CreatedBy int
}

func (p *ProjectService) CreateComment(ctx context.Context, opt *ProjectCommentOpt) (int64, error) {
	comment := &model.ProjectTaskComment{
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

func (p *ProjectService) Comments(ctx context.Context, TaskId int64) ([]*model.ProjectTaskComment, error) {
	query := p.Db().WithContext(ctx).Table("project_task_comments")
	query.Where("task_id = ?", TaskId)

	var items []*model.ProjectTaskComment
	if err := query.Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}
