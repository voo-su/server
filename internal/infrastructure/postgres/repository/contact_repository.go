package repository

import (
	"context"
	"gorm.io/gorm"
	"strconv"
	"voo.su/internal/constant"
	postgresModel "voo.su/internal/infrastructure/postgres/model"
	redisRepo "voo.su/internal/infrastructure/redis/repository"
	"voo.su/pkg/repo"
)

type ContactRepository struct {
	repo.Repo[postgresModel.Contact]
	ContactRemarkCacheRepo *redisRepo.ContactRemarkCacheRepository
	RelationCacheRepo      *redisRepo.RelationCacheRepository
}

func NewContactRepository(
	db *gorm.DB,
	contactRemarkCacheRepo *redisRepo.ContactRemarkCacheRepository,
	relationCacheRepo *redisRepo.RelationCacheRepository,
) *ContactRepository {
	return &ContactRepository{
		Repo:                   repo.NewRepo[postgresModel.Contact](db),
		ContactRemarkCacheRepo: contactRemarkCacheRepo,
		RelationCacheRepo:      relationCacheRepo,
	}
}

func (c *ContactRepository) Remarks(ctx context.Context, uid int, fids []int) (map[int]string, error) {
	if !c.ContactRemarkCacheRepo.Exist(ctx, uid) {
		_ = c.LoadContactCache(ctx, uid)
	}
	return c.ContactRemarkCacheRepo.MGet(ctx, uid, fids)
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

func (c *ContactRepository) GetFriendRemark(ctx context.Context, uid int, friendId int) string {
	if c.ContactRemarkCacheRepo.Exist(ctx, uid) {
		return c.ContactRemarkCacheRepo.Get(ctx, uid, friendId)
	}

	var remark string
	c.Repo.Model(ctx).Where("user_id = ? AND friend_id = ?", uid, friendId).Pluck("remark", &remark)

	return remark
}

func (c *ContactRepository) SetFriendRemark(ctx context.Context, uid int, friendId int, remark string) error {
	return c.ContactRemarkCacheRepo.Set(ctx, uid, friendId, remark)
}

func (c *ContactRepository) LoadContactCache(ctx context.Context, uid int) error {
	all, err := c.Repo.FindAll(ctx, func(db *gorm.DB) {
		db.Select("friend_id,remark").Where("user_id = ? AND status = ?", uid, constant.ContactStatusNormal)
	})
	if err != nil {
		return err
	}

	items := make(map[string]any)
	for _, value := range all {
		if len(value.Remark) > 0 {
			items[strconv.Itoa(value.FriendId)] = value.Remark
		}
	}

	return c.ContactRemarkCacheRepo.MSet(ctx, uid, items)
}
