package repo

import (
	"context"
	"gorm.io/gorm"
	"voo.su/internal/domain/entity"
	"voo.su/internal/repository/model"
	"voo.su/pkg/repo"
)

type GroupChatAds struct {
	repo.Repo[model.GroupChatAds]
}

func NewGroupChatAds(db *gorm.DB) *GroupChatAds {
	return &GroupChatAds{Repo: repo.NewRepo[model.GroupChatAds](db)}
}

func (g *GroupChatAds) GetListAll(ctx context.Context, groupId int) ([]*entity.SearchAdsItem, error) {
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
	query := g.Repo.Db.WithContext(ctx).Table("group_chat_ads")
	query.Joins("LEFT JOIN users on users.id = group_chat_ads.creator_id")
	query.Where("group_chat_ads.group_id = ? and group_chat_ads.is_delete = ?", groupId, 0)
	query.Order("group_chat_ads.is_top desc")
	query.Order("group_chat_ads.created_at desc")

	var items []*entity.SearchAdsItem
	if err := query.Select(fields).Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (g *GroupChatAds) GetLatestAds(ctx context.Context, groupId int) (*model.GroupChatAds, error) {
	var info model.GroupChatAds
	err := g.Repo.Db.WithContext(ctx).Last(&info, "group_id = ? and is_delete = ?", groupId, 0).Error
	if err != nil {
		return nil, err
	}

	return &info, nil
}
