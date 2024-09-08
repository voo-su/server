package service

import (
	"context"
	"time"
	"voo.su/internal/repository/model"
)

type ProjectCommentOpt struct {
	TaskId      int64
	CommentText string
	CreatedBy   int
}

func (p *ProjectService) CreateComment(ctx context.Context, opt *ProjectCommentOpt) (int64, error) {
	comment := &model.ProjectTaskComment{
		TaskId:    int(opt.TaskId),
		CreatedBy: opt.CreatedBy,
		CreatedAt: time.Now(),
	}

	if err := p.ProjectTaskComment.Create(ctx, comment); err != nil {
		return int64(comment.Id), err
	}

	return int64(comment.Id), nil
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
