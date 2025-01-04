package queue

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type EmailHandle struct {
	Redis *redis.Client
}

func (e *EmailHandle) Handle(ctx context.Context) error {
	return nil
}
