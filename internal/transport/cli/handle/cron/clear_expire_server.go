package cron

import (
	"context"
	"voo.su/internal/repository/cache"
)

type ClearExpireServer struct {
	storage *cache.ServerStorage
}

func NewClearExpireServer(storage *cache.ServerStorage) *ClearExpireServer {
	return &ClearExpireServer{storage: storage}
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
	for _, sid := range c.storage.All(ctx, 2) {
		_ = c.storage.Del(ctx, sid)
		_ = c.storage.SetExpireServer(ctx, sid)
	}

	return nil
}
