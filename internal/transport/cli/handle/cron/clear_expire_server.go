package cron

import (
	"context"
	"voo.su/internal/repository/cache"
)

type ClearExpireServer struct {
	Storage *cache.ServerStorage
}

func NewClearExpireServer(storage *cache.ServerStorage) *ClearExpireServer {
	return &ClearExpireServer{Storage: storage}
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
	for _, sid := range c.Storage.All(ctx, 2) {
		_ = c.Storage.Del(ctx, sid)
		_ = c.Storage.SetExpireServer(ctx, sid)
	}

	return nil
}
