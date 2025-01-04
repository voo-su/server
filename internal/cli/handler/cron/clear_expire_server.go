package cron

import (
	"context"
	redisRepo "voo.su/internal/infrastructure/redis/repository"
)

type ClearExpireServer struct {
	ServerCacheRepo *redisRepo.ServerCacheRepository
}

func NewClearExpireServer(serverCacheRepo *redisRepo.ServerCacheRepository) *ClearExpireServer {
	return &ClearExpireServer{ServerCacheRepo: serverCacheRepo}
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
	for _, sid := range c.ServerCacheRepo.All(ctx, 2) {
		_ = c.ServerCacheRepo.Del(ctx, sid)
		_ = c.ServerCacheRepo.SetExpireServer(ctx, sid)
	}

	return nil
}
