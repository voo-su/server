package cron

import (
	"context"
	"fmt"
	"voo.su/internal/repository/cache"
)

type ClearWsCache struct {
	Storage *cache.ServerStorage
}

func NewClearWsCache(storage *cache.ServerStorage) *ClearWsCache {
	return &ClearWsCache{Storage: storage}
}

func (c *ClearWsCache) Name() string {
	return "clear.ws.cache"
}

func (c *ClearWsCache) Spec() string {
	return "*/30 * * * *"
}

func (c *ClearWsCache) Enable() bool {
	return true
}

func (c *ClearWsCache) Handle(ctx context.Context) error {
	for _, sid := range c.Storage.GetExpireServerAll(ctx) {
		c.clear(ctx, sid)
	}
	return nil
}

func (c *ClearWsCache) clear(ctx context.Context, sid string) {
	var cursor uint64
	for {
		var keys []string
		var err error
		keys, cursor, err = c.Storage.Redis().Scan(ctx, cursor, fmt.Sprintf("ws:%s:*", sid), 200).Result()
		if err != nil {
			return
		}

		c.Storage.Redis().Del(ctx, keys...)
		if cursor == 0 {
			_ = c.Storage.DelExpireServer(ctx, sid)
			break
		}
	}
}
