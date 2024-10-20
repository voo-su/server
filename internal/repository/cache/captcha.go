package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type CaptchaStorage struct {
	Redis *redis.Client
}

func NewCaptchaStorage(redis *redis.Client) *CaptchaStorage {
	return &CaptchaStorage{Redis: redis}
}

func (c *CaptchaStorage) Set(id string, value string) error {
	return c.Redis.SetEx(context.TODO(), c.name(id), value, 3*time.Minute).Err()
}

func (c *CaptchaStorage) Get(id string, clear bool) string {
	value := c.Redis.Get(context.TODO(), c.name(id)).Val()
	if clear && len(value) > 0 {
		c.Redis.Del(context.TODO(), c.name(id))
	}

	return value
}

func (c *CaptchaStorage) Verify(id, answer string, clear bool) bool {
	value := c.Redis.Get(context.TODO(), c.name(id)).Val()
	if clear && len(value) > 0 {
		c.Redis.Del(context.TODO(), c.name(id))
	}

	return value == answer
}

func (c *CaptchaStorage) name(id string) string {
	return fmt.Sprintf("im:auth:captcha:%s", id)
}
