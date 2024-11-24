package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
)

type UnreadStorage struct {
	Rds *redis.Client
}

func NewUnreadStorage(rds *redis.Client) *UnreadStorage {
	return &UnreadStorage{rds}
}

func (u *UnreadStorage) Incr(ctx context.Context, mode, sender, receive int) {
	u.Rds.HIncrBy(ctx, u.name(receive), fmt.Sprintf("%d_%d", mode, sender), 1)
}

func (u *UnreadStorage) PipeIncr(ctx context.Context, pipe redis.Pipeliner, mode, sender, receive int) {
	pipe.HIncrBy(ctx, u.name(receive), fmt.Sprintf("%d_%d", mode, sender), 1)
}

func (u *UnreadStorage) Get(ctx context.Context, mode, sender, receive int) int {
	val, _ := u.Rds.HGet(ctx, u.name(receive), fmt.Sprintf("%d_%d", mode, sender)).Int()
	return val
}

func (u *UnreadStorage) Del(ctx context.Context, mode, sender, receive int) {
	u.Rds.HDel(ctx, u.name(receive), fmt.Sprintf("%d_%d", mode, sender))
}

func (u *UnreadStorage) Reset(ctx context.Context, mode, sender, receive int) {
	u.Rds.HSet(ctx, u.name(receive), fmt.Sprintf("%d_%d", mode, sender), 0)
}

func (u *UnreadStorage) All(ctx context.Context, receive int) map[string]int {
	items := make(map[string]int)
	for k, v := range u.Rds.HGetAll(ctx, u.name(receive)).Val() {
		items[k], _ = strconv.Atoi(v)
	}
	return items
}

func (u *UnreadStorage) name(receive int) string {
	return fmt.Sprintf("im:message:unread:uid_%d", receive)
}
