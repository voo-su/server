package queue

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type LoginHandle struct {
	rds *redis.Client
}

func (e *LoginHandle) Handle(ctx context.Context) error {
	return nil
}
