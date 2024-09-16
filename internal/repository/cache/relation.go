package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type Relation struct {
	Redis *redis.Client
}

func NewRelation(redis *redis.Client) *Relation {
	return &Relation{Redis: redis}
}

func (r *Relation) IsContactRelation(ctx context.Context, uid, uid2 int) error {
	return r.Redis.Get(ctx, r.keyContactRelation(uid, uid2)).Err()
}

func (r *Relation) SetContactRelation(ctx context.Context, uid, uid2 int) {
	r.Redis.SetEx(ctx, r.keyContactRelation(uid, uid2), "1", time.Hour*1)
}

func (r *Relation) DelContactRelation(ctx context.Context, uid, uid2 int) {
	r.Redis.Del(ctx, r.keyContactRelation(uid, uid2))
}

func (r *Relation) IsGroupRelation(ctx context.Context, uid, gid int) error {
	return r.Redis.Get(ctx, r.keyGroupRelation(uid, gid)).Err()
}

func (r *Relation) SetGroupRelation(ctx context.Context, uid, gid int) {
	r.Redis.SetEx(ctx, r.keyGroupRelation(uid, gid), "1", time.Hour*1)
}

func (r *Relation) DelGroupRelation(ctx context.Context, uid, gid int) {
	r.Redis.Del(ctx, r.keyGroupRelation(uid, gid))
}

func (r *Relation) BatchDelGroupRelation(ctx context.Context, uids []int, gid int) {
	for _, uid := range uids {
		r.DelGroupRelation(ctx, uid, gid)
	}
}

func (r *Relation) keyContactRelation(uid, uid2 int) string {
	if uid2 < uid {
		uid, uid2 = uid2, uid
	}

	return fmt.Sprintf("im:contact:relation:%d_%d", uid, uid2)
}

func (r *Relation) keyGroupRelation(uid, gid int) string {
	return fmt.Sprintf("im:contact:relation:%d_%d", uid, gid)
}
