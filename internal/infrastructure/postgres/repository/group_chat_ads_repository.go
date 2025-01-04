package repository

import (
	"context"
	"gorm.io/gorm"
	"voo.su/internal/domain/entity"
	"voo.su/internal/infrastructure/postgres/model"
	"voo.su/pkg/gormutil"
)

type GroupChatAdsRepository struct {
	gormutil.Repo[model.GroupChatAds]
}

func NewGroupChatAdsRepository(db *gorm.DB) *GroupChatAdsRepository {
	return &GroupChatAdsRepository{Repo: gormutil.NewRepo[model.GroupChatAds](db)}
}

func (g *GroupChatAdsRepository) GetListAll(ctx context.Context, groupId int) ([]*entity.SearchAdsItem, error) {
	fields := []string{
		"group_chat_ads.id",
		"group_chat_ads.creator_id",
		"group_chat_ads.title",
		"group_chat_ads.content",
		"group_chat_ads.is_top",
		"group_chat_ads.is_confirm",
		"group_chat_ads.confirm_users",
		"group_chat_ads.created_at",
		"group_chat_ads.updated_at",
		"users.avatar",
		"users.username",
	}
	query := g.Repo.Db.WithContext(ctx).
		Select(fields).
		Table("group_chat_ads").
		Joins("LEFT JOIN users on users.id = group_chat_ads.creator_id").
		Where("group_chat_ads.group_id = ? AND group_chat_ads.is_delete = ?", groupId, 0).
		Order("group_chat_ads.is_top desc").
		Order("group_chat_ads.created_at desc")

	var items []*entity.SearchAdsItem
	if err := query.Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (g *GroupChatAdsRepository) GetLatestAds(ctx context.Context, groupId int) (*model.GroupChatAds, error) {
	var info model.GroupChatAds
	if err := g.Repo.Db.WithContext(ctx).
		Last(&info, "group_id = ? AND is_delete = ?", groupId, 0).
		Error; err != nil {
		return nil, err
	}

	return &info, nil
}
