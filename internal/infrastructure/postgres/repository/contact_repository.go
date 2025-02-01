package repository

import (
	"context"
	"gorm.io/gorm"
	"voo.su/internal/constant"
	postgresModel "voo.su/internal/infrastructure/postgres/model"
	redisRepo "voo.su/internal/infrastructure/redis/repository"
	"voo.su/pkg/gormutil"
)

type ContactRepository struct {
	gormutil.Repo[postgresModel.Contact]
	RelationCacheRepo *redisRepo.RelationCacheRepository
}

func NewContactRepository(
	db *gorm.DB,
	relationCacheRepo *redisRepo.RelationCacheRepository,
) *ContactRepository {
	return &ContactRepository{
		Repo:              gormutil.NewRepo[postgresModel.Contact](db),
		RelationCacheRepo: relationCacheRepo,
	}
}

func (c *ContactRepository) IsFriend(ctx context.Context, uid int, friendId int, cache bool) bool {
	if cache && c.RelationCacheRepo.IsContactRelation(ctx, uid, friendId) == nil {
		return true
	}
	count, err := c.Repo.QueryCount(ctx, "((user_id = ? AND friend_id = ?) OR (user_id = ? AND friend_id = ?)) AND status = ?", uid, friendId, friendId, uid, constant.ContactStatusNormal)
	if err != nil {
		return false
	}

	if count == 2 {
		c.RelationCacheRepo.SetContactRelation(ctx, uid, friendId)
	} else {
		c.RelationCacheRepo.DelContactRelation(ctx, uid, friendId)
	}

	return count == 2
}
