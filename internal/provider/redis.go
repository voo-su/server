package provider

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"voo.su/internal/config"
)

func NewRedisClient(conf *config.Config) *redis.Client {
	client := redis.NewClient(conf.Redis.Options())
	if _, err := client.Ping(context.TODO()).Result(); err != nil {
		panic(fmt.Errorf("ошибка клиента redis: %s", err))
	}

	return client
}
