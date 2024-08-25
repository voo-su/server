package service

import (
	"context"
	"time"
	"voo.su/internal/repository/model"
	"voo.su/internal/repository/repo"
	"voo.su/pkg/timeutil"
)

type GroupChatAdsService struct {
	*repo.Source
	Ads *repo.GroupChatAds
}

func NewGroupChatAdsService(
	source *repo.Source,
	ads *repo.GroupChatAds,
) *GroupChatAdsService {
	return &GroupChatAdsService{
		Source: source,
		Ads:    ads,
	}
}

type GroupChatAdsEditOpt struct {
	UserId    int
	GroupId   int
	AdsId     int
	Title     string
	Content   string
	IsTop     int
	IsConfirm int
}

func (s *GroupChatAdsService) Create(ctx context.Context, opt *GroupChatAdsEditOpt) error {
	return s.Ads.Create(ctx, &model.GroupChatAds{
		GroupId:      opt.GroupId,
		CreatorId:    opt.UserId,
		Title:        opt.Title,
		Content:      opt.Content,
		IsTop:        opt.IsTop,
		IsConfirm:    opt.IsConfirm,
		ConfirmUsers: "{}",
	})
}

func (s *GroupChatAdsService) Update(ctx context.Context, opt *GroupChatAdsEditOpt) error {
	_, err := s.Ads.UpdateWhere(ctx, map[string]any{
		"title":      opt.Title,
		"content":    opt.Content,
		"is_top":     opt.IsTop,
		"is_confirm": opt.IsConfirm,
		"updated_at": time.Now(),
	}, "id = ? and group_id = ?", opt.AdsId, opt.GroupId)

	return err
}

func (s *GroupChatAdsService) Delete(ctx context.Context, groupId, id int) error {
	_, err := s.Ads.UpdateWhere(ctx, map[string]any{
		"is_delete":  1,
		"deleted_at": timeutil.DateTime(),
		"updated_at": time.Now(),
	}, "id = ? and group_id = ?", id, groupId)

	return err
}
