package usecase

import (
	"context"
	"time"
	"voo.su/internal/infrastructure"
	postgresModel "voo.su/internal/infrastructure/postgres/model"
	postgresRepo "voo.su/internal/infrastructure/postgres/repository"
	"voo.su/pkg/locale"
	"voo.su/pkg/timeutil"
)

type GroupChatAdsUseCase struct {
	Locale           locale.ILocale
	Source           *infrastructure.Source
	GroupChatAdsRepo *postgresRepo.GroupChatAdsRepository
}

func NewGroupChatAdsUseCase(
	locale locale.ILocale,
	source *infrastructure.Source,
	groupChatAdsRepo *postgresRepo.GroupChatAdsRepository,
) *GroupChatAdsUseCase {
	return &GroupChatAdsUseCase{
		Locale:           locale,
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
	return g.GroupChatAdsRepo.Create(ctx, &postgresModel.GroupChatAds{
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
	}, "id = ? AND group_id = ?", opt.AdsId, opt.GroupId)

	return err
}

func (g *GroupChatAdsUseCase) Delete(ctx context.Context, groupId, id int) error {
	_, err := g.GroupChatAdsRepo.UpdateWhere(ctx, map[string]any{
		"is_delete":  1,
		"deleted_at": timeutil.DateTime(),
		"updated_at": time.Now(),
	}, "id = ? AND group_id = ?", id, groupId)

	return err
}
