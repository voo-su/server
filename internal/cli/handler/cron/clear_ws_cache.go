// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package cron

import (
	"context"
	"fmt"
	redisRepo "voo.su/internal/infrastructure/redis/repository"
)

type ClearWsCache struct {
	ServerCacheRepo *redisRepo.ServerCacheRepository
}

func NewClearWsCache(serverCacheRepo *redisRepo.ServerCacheRepository) *ClearWsCache {
	return &ClearWsCache{ServerCacheRepo: serverCacheRepo}
}

func (c *ClearWsCache) Name() string {
	return "clear.ws.redis"
}

func (c *ClearWsCache) Spec() string {
	return "*/30 * * * *"
}

func (c *ClearWsCache) Enable() bool {
	return true
}

func (c *ClearWsCache) Handle(ctx context.Context) error {
	for _, sid := range c.ServerCacheRepo.GetExpireServerAll(ctx) {
		c.clear(ctx, sid)
	}
	return nil
}

func (c *ClearWsCache) clear(ctx context.Context, sid string) {
	var cursor uint64
	for {
		var keys []string
		var err error
		keys, cursor, err = c.ServerCacheRepo.Redis().Scan(ctx, cursor, fmt.Sprintf("ws:%s:*", sid), 200).Result()
		if err != nil {
			return
		}

		c.ServerCacheRepo.Redis().Del(ctx, keys...)
		if cursor == 0 {
			_ = c.ServerCacheRepo.DelExpireServer(ctx, sid)
			break
		}
	}
}
