// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package repository

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
)

type UnreadCacheRepository struct {
	Rds *redis.Client
}

func NewUnreadCacheRepository(rds *redis.Client) *UnreadCacheRepository {
	return &UnreadCacheRepository{rds}
}

func (u *UnreadCacheRepository) Incr(ctx context.Context, mode, sender, receive int) {
	u.Rds.HIncrBy(ctx, u.name(receive), fmt.Sprintf("%d_%d", mode, sender), 1)
}

func (u *UnreadCacheRepository) PipeIncr(ctx context.Context, pipe redis.Pipeliner, mode, sender, receive int) {
	pipe.HIncrBy(ctx, u.name(receive), fmt.Sprintf("%d_%d", mode, sender), 1)
}

func (u *UnreadCacheRepository) Get(ctx context.Context, mode, sender, receive int) int {
	val, _ := u.Rds.HGet(ctx, u.name(receive), fmt.Sprintf("%d_%d", mode, sender)).Int()
	return val
}

func (u *UnreadCacheRepository) Del(ctx context.Context, mode, sender, receive int) {
	u.Rds.HDel(ctx, u.name(receive), fmt.Sprintf("%d_%d", mode, sender))
}

func (u *UnreadCacheRepository) Reset(ctx context.Context, mode, sender, receive int) {
	u.Rds.HSet(ctx, u.name(receive), fmt.Sprintf("%d_%d", mode, sender), 0)
}

func (u *UnreadCacheRepository) All(ctx context.Context, receive int) map[string]int {
	items := make(map[string]int)
	for k, v := range u.Rds.HGetAll(ctx, u.name(receive)).Val() {
		items[k], _ = strconv.Atoi(v)
	}
	return items
}

func (u *UnreadCacheRepository) name(receive int) string {
	return fmt.Sprintf("im:message:unread:uid_%d", receive)
}
