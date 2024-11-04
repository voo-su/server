package usecase

import (
	"context"
	"time"
	"voo.su/internal/repository/model"
	"voo.su/internal/repository/repo"
	"voo.su/pkg/timeutil"
)

type GroupChatAdsUseCase struct {
	*repo.Source
	GroupChatAdsRepo *repo.GroupChatAds
}

func NewGroupChatAdsUseCase(
	source *repo.Source,
	groupChatAdsRepo *repo.GroupChatAds,
) *GroupChatAdsUseCase {
	return &GroupChatAdsUseCase{
		Source:           source,
		GroupChatAdsRepo: groupChatAdsRepo,
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

func (g *GroupChatAdsUseCase) Create(ctx context.Context, opt *GroupChatAdsEditOpt) error {
	return g.GroupChatAdsRepo.Create(ctx, &model.GroupChatAds{
		GroupId:      opt.GroupId,
		CreatorId:    opt.UserId,
		Title:        opt.Title,
		Content:      opt.Content,
		IsTop:        opt.IsTop,
		IsConfirm:    opt.IsConfirm,
		ConfirmUsers: "{}",
	})
}

func (g *GroupChatAdsUseCase) Update(ctx context.Context, opt *GroupChatAdsEditOpt) error {
	_, err := g.GroupChatAdsRepo.UpdateWhere(ctx, map[string]any{
		"title":      opt.Title,
		"content":    opt.Content,
		"is_top":     opt.IsTop,
		"is_confirm": opt.IsConfirm,
		"updated_at": time.Now(),
	}, "id = ? and group_id = ?", opt.AdsId, opt.GroupId)

	return err
}

func (g *GroupChatAdsUseCase) Delete(ctx context.Context, groupId, id int) error {
	_, err := g.GroupChatAdsRepo.UpdateWhere(ctx, map[string]any{
		"is_delete":  1,
		"deleted_at": timeutil.DateTime(),
		"updated_at": time.Now(),
	}, "id = ? and group_id = ?", id, groupId)

	return err
}
