package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
)

type UnreadStorage struct {
	redis *redis.Client
}

func NewUnreadStorage(rds *redis.Client) *UnreadStorage {
	return &UnreadStorage{rds}
}

func (u *UnreadStorage) Incr(ctx context.Context, mode, sender, receive int) {
	u.redis.HIncrBy(ctx, u.name(receive), fmt.Sprintf("%d_%d", mode, sender), 1)
}

func (u *UnreadStorage) PipeIncr(ctx context.Context, pipe redis.Pipeliner, mode, sender, receive int) {
	pipe.HIncrBy(ctx, u.name(receive), fmt.Sprintf("%d_%d", mode, sender), 1)
}

func (u *UnreadStorage) Get(ctx context.Context, mode, sender, receive int) int {
	val, _ := u.redis.HGet(ctx, u.name(receive), fmt.Sprintf("%d_%d", mode, sender)).Int()
	return val
}

func (u *UnreadStorage) Del(ctx context.Context, mode, sender, receive int) {
	u.redis.HDel(ctx, u.name(receive), fmt.Sprintf("%d_%d", mode, sender))
}

func (u *UnreadStorage) Reset(ctx context.Context, mode, sender, receive int) {
	u.redis.HSet(ctx, u.name(receive), fmt.Sprintf("%d_%d", mode, sender), 0)
}

func (u *UnreadStorage) All(ctx context.Context, receive int) map[string]int {
	items := make(map[string]int)
	for k, v := range u.redis.HGetAll(ctx, u.name(receive)).Val() {
		items[k], _ = strconv.Atoi(v)
	}
	return items
}

func (u *UnreadStorage) name(receive int) string {
	return fmt.Sprintf("im:message:unread:uid_%d", receive)
}
