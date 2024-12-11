package repository

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type CaptchaCacheRepository struct {
	Rds *redis.Client
}

func NewCaptchaCacheRepository(redis *redis.Client) *CaptchaCacheRepository {
	return &CaptchaCacheRepository{Rds: redis}
}

func (c *CaptchaCacheRepository) Set(id string, value string) error {
	return c.Rds.SetEx(context.TODO(), c.name(id), value, 3*time.Minute).Err()
}

func (c *CaptchaCacheRepository) Get(id string, clear bool) string {
	value := c.Rds.Get(context.TODO(), c.name(id)).Val()
	if clear && len(value) > 0 {
		c.Rds.Del(context.TODO(), c.name(id))
	}

	return value
}

func (c *CaptchaCacheRepository) Verify(id, answer string, clear bool) bool {
	value := c.Rds.Get(context.TODO(), c.name(id)).Val()
	if clear && len(value) > 0 {
		c.Rds.Del(context.TODO(), c.name(id))
	}

	return value == answer
}

func (c *CaptchaCacheRepository) name(id string) string {
	return fmt.Sprintf("im:auth:captcha:%s", id)
}
