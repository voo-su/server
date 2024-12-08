package cron

import (
	"context"
	"voo.su/internal/repository/cache"
)

type ClearExpireServer struct {
	ServerCache *cache.ServerCache
}

func NewClearExpireServer(serverCache *cache.ServerCache) *ClearExpireServer {
	return &ClearExpireServer{ServerCache: serverCache}
}

func (c *ClearExpireServer) Name() string {
	return "clear.expire.server"
}

func (c *ClearExpireServer) Spec() string {
	return "*/5 * * * *"
}

func (c *ClearExpireServer) Enable() bool {
	return true
}

func (c *ClearExpireServer) Handle(ctx context.Context) error {
	for _, sid := range c.ServerCache.All(ctx, 2) {
		_ = c.ServerCache.Del(ctx, sid)
		_ = c.ServerCache.SetExpireServer(ctx, sid)
	}

	return nil
}
