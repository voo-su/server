package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisLock struct {
	Rds *redis.Client
}

func NewRedisLock(rds *redis.Client) *RedisLock {
	return &RedisLock{rds}
}

func (r *RedisLock) Lock(ctx context.Context, name string, expire int) bool {
	return r.Rds.SetNX(ctx, r.name(name), 1, time.Duration(expire)*time.Second).Val()
}

func (r *RedisLock) UnLock(ctx context.Context, name string) bool {
	script := `
	if redis.call("GET", KEYS[1]) == ARGV[1] then
		return redis.call("DEL", KEYS[1])
	else
		return false
	end`
	return r.Rds.Eval(ctx, script, []string{r.name(name)}, 1).Err() == nil
}

func (r *RedisLock) name(name string) string {
	return fmt.Sprintf("im:lock:%s", name)
}
