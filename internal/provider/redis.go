// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package provider

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"voo.su/internal/config"
	"voo.su/pkg/locale"
)

func NewRedisClient(conf *config.Config, locale locale.ILocale) *redis.Client {
	client := redis.NewClient(conf.Redis.Options())
	if _, err := client.Ping(context.TODO()).Result(); err != nil {
		panic(fmt.Errorf(locale.Localize("connection_error"), "Redis", err))
	}

	return client
}
