// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

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
