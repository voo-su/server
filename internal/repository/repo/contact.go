package repo

import (
	"context"
	"gorm.io/gorm"
	"strconv"
	"voo.su/internal/repository/cache"
	"voo.su/internal/repository/model"
	"voo.su/pkg/core"
)

type Contact struct {
	core.Repo[model.Contact]
	Cache    *cache.ContactRemark
	Relation *cache.Relation
}

func NewContact(
	db *gorm.DB,
	cache *cache.ContactRemark,
	relation *cache.Relation,
) *Contact {
	return &Contact{
		Repo:     core.NewRepo[model.Contact](db),
		Cache:    cache,
		Relation: relation,
	}
}

func (c *Contact) Remarks(ctx context.Context, uid int, fids []int) (map[int]string, error) {
	if !c.Cache.Exist(ctx, uid) {
		_ = c.LoadContactCache(ctx, uid)
	}
	return c.Cache.MGet(ctx, uid, fids)
}

func (c *Contact) IsFriend(ctx context.Context, uid int, friendId int, cache bool) bool {
	if cache && c.Relation.IsContactRelation(ctx, uid, friendId) == nil {
		return true
	}
	count, err := c.Repo.QueryCount(ctx, "((user_id = ? and friend_id = ?) or (user_id = ? and friend_id = ?)) and status = ?", uid, friendId, friendId, uid, model.ContactStatusNormal)
	if err != nil {
		return false
	}

	if count == 2 {
		c.Relation.SetContactRelation(ctx, uid, friendId)
	} else {
		c.Relation.DelContactRelation(ctx, uid, friendId)
	}

	return count == 2
}

func (c *Contact) GetFriendRemark(ctx context.Context, uid int, friendId int) string {
	if c.Cache.Exist(ctx, uid) {
		return c.Cache.Get(ctx, uid, friendId)
	}

	var remark string
	c.Repo.Model(ctx).Where("user_id = ? and friend_id = ?", uid, friendId).Pluck("remark", &remark)

	return remark
}

func (c *Contact) SetFriendRemark(ctx context.Context, uid int, friendId int, remark string) error {
	return c.Cache.Set(ctx, uid, friendId, remark)
}

func (c *Contact) LoadContactCache(ctx context.Context, uid int) error {
	all, err := c.Repo.FindAll(ctx, func(db *gorm.DB) {
		db.Select("friend_id,remark").Where("user_id = ? and status = ?", uid, model.ContactStatusNormal)
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

	return c.Cache.MSet(ctx, uid, items)
}
