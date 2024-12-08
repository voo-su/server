package repo

import (
	"context"
	"gorm.io/gorm"
	"strconv"
	"voo.su/internal/constant"
	"voo.su/internal/repository/cache"
	"voo.su/internal/repository/model"
	"voo.su/pkg/repo"
)

type Contact struct {
	repo.Repo[model.Contact]
	ContactRemarkCache *cache.ContactRemarkCache
	RelationCache      *cache.RelationCache
}

func NewContact(
	db *gorm.DB,
	contactRemarkCache *cache.ContactRemarkCache,
	relation *cache.RelationCache,
) *Contact {
	return &Contact{
		Repo:               repo.NewRepo[model.Contact](db),
		ContactRemarkCache: contactRemarkCache,
		RelationCache:      relation,
	}
}

func (c *Contact) Remarks(ctx context.Context, uid int, fids []int) (map[int]string, error) {
	if !c.ContactRemarkCache.Exist(ctx, uid) {
		_ = c.LoadContactCache(ctx, uid)
	}
	return c.ContactRemarkCache.MGet(ctx, uid, fids)
}

func (c *Contact) IsFriend(ctx context.Context, uid int, friendId int, cache bool) bool {
	if cache && c.RelationCache.IsContactRelation(ctx, uid, friendId) == nil {
		return true
	}
	count, err := c.Repo.QueryCount(ctx, "((user_id = ? AND friend_id = ?) or (user_id = ? AND friend_id = ?)) AND status = ?", uid, friendId, friendId, uid, constant.ContactStatusNormal)
	if err != nil {
		return false
	}

	if count == 2 {
		c.RelationCache.SetContactRelation(ctx, uid, friendId)
	} else {
		c.RelationCache.DelContactRelation(ctx, uid, friendId)
	}

	return count == 2
}

func (c *Contact) GetFriendRemark(ctx context.Context, uid int, friendId int) string {
	if c.ContactRemarkCache.Exist(ctx, uid) {
		return c.ContactRemarkCache.Get(ctx, uid, friendId)
	}

	var remark string
	c.Repo.Model(ctx).Where("user_id = ? AND friend_id = ?", uid, friendId).Pluck("remark", &remark)

	return remark
}

func (c *Contact) SetFriendRemark(ctx context.Context, uid int, friendId int, remark string) error {
	return c.ContactRemarkCache.Set(ctx, uid, friendId, remark)
}

func (c *Contact) LoadContactCache(ctx context.Context, uid int) error {
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

	return c.ContactRemarkCache.MSet(ctx, uid, items)
}
