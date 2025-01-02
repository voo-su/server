// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package repository

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisLockCacheRepository struct {
	Rds *redis.Client
}

func NewRedisLockCacheRepository(rds *redis.Client) *RedisLockCacheRepository {
	return &RedisLockCacheRepository{
		Rds: rds,
	}
}

func (r *RedisLockCacheRepository) Lock(ctx context.Context, name string, expire int) bool {
	return r.Rds.SetNX(ctx, r.name(name), 1, time.Duration(expire)*time.Second).Val()
}

func (r *RedisLockCacheRepository) UnLock(ctx context.Context, name string) bool {
	script := `
	if redis.call("GET", KEYS[1]) == ARGV[1] then
		return redis.call("DEL", KEYS[1])
	else
		return false
	end`
	return r.Rds.Eval(ctx, script, []string{r.name(name)}, 1).Err() == nil
}

func (r *RedisLockCacheRepository) name(name string) string {
	return fmt.Sprintf("im:lock:%s", name)
}
