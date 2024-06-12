package service

import (
	"context"
	"time"
	"voo.su/internal/repository/model"
	"voo.su/internal/repository/repo"
	"voo.su/pkg/timeutil"
)

type GroupChatNoticeService struct {
	*repo.Source
	Notice *repo.GroupChatNotice
}

func NewGroupChatNoticeService(
	source *repo.Source,
	notice *repo.GroupChatNotice,
) *GroupChatNoticeService {
	return &GroupChatNoticeService{
		Source: source,
		Notice: notice,
	}
}

type GroupChatNoticeEditOpt struct {
	UserId    int
	GroupId   int
	NoticeId  int
	Title     string
	Content   string
	IsTop     int
	IsConfirm int
}

func (s *GroupChatNoticeService) Create(ctx context.Context, opt *GroupChatNoticeEditOpt) error {
	return s.Notice.Create(ctx, &model.GroupChatNotice{
		GroupId:      opt.GroupId,
		CreatorId:    opt.UserId,
		Title:        opt.Title,
		Content:      opt.Content,
		IsTop:        opt.IsTop,
		IsConfirm:    opt.IsConfirm,
		ConfirmUsers: "{}",
	})
}

func (s *GroupChatNoticeService) Update(ctx context.Context, opt *GroupChatNoticeEditOpt) error {
	_, err := s.Notice.UpdateWhere(ctx, map[string]any{
		"title":      opt.Title,
		"content":    opt.Content,
		"is_top":     opt.IsTop,
		"is_confirm": opt.IsConfirm,
		"updated_at": time.Now(),
	}, "id = ? and group_id = ?", opt.NoticeId, opt.GroupId)

	return err
}

func (s *GroupChatNoticeService) Delete(ctx context.Context, groupId, noticeId int) error {
	_, err := s.Notice.UpdateWhere(ctx, map[string]any{
		"is_delete":  1,
		"deleted_at": timeutil.DateTime(),
		"updated_at": time.Now(),
	}, "id = ? and group_id = ?", noticeId, groupId)

	return err
}
