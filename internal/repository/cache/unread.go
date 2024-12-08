package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
)

type UnreadCache struct {
	Rds *redis.Client
}

func NewUnreadCache(rds *redis.Client) *UnreadCache {
	return &UnreadCache{rds}
}

func (u *UnreadCache) Incr(ctx context.Context, mode, sender, receive int) {
	u.Rds.HIncrBy(ctx, u.name(receive), fmt.Sprintf("%d_%d", mode, sender), 1)
}

func (u *UnreadCache) PipeIncr(ctx context.Context, pipe redis.Pipeliner, mode, sender, receive int) {
	pipe.HIncrBy(ctx, u.name(receive), fmt.Sprintf("%d_%d", mode, sender), 1)
}

func (u *UnreadCache) Get(ctx context.Context, mode, sender, receive int) int {
	val, _ := u.Rds.HGet(ctx, u.name(receive), fmt.Sprintf("%d_%d", mode, sender)).Int()
	return val
}

func (u *UnreadCache) Del(ctx context.Context, mode, sender, receive int) {
	u.Rds.HDel(ctx, u.name(receive), fmt.Sprintf("%d_%d", mode, sender))
}

func (u *UnreadCache) Reset(ctx context.Context, mode, sender, receive int) {
	u.Rds.HSet(ctx, u.name(receive), fmt.Sprintf("%d_%d", mode, sender), 0)
}

func (u *UnreadCache) All(ctx context.Context, receive int) map[string]int {
	items := make(map[string]int)
	for k, v := range u.Rds.HGetAll(ctx, u.name(receive)).Val() {
		items[k], _ = strconv.Atoi(v)
	}
	return items
}

func (u *UnreadCache) name(receive int) string {
	return fmt.Sprintf("im:message:unread:uid_%d", receive)
}
