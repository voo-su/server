package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

type ContactRemarkCache struct {
	Rds *redis.Client
}

func NewContactRemarkCache(redis *redis.Client) *ContactRemarkCache {
	return &ContactRemarkCache{Rds: redis}
}

func (c *ContactRemarkCache) Get(ctx context.Context, uid int, fid int) string {
	return c.Rds.HGet(ctx, c.name(uid), fmt.Sprintf("%d", fid)).Val()
}

func (c *ContactRemarkCache) MGet(ctx context.Context, uid int, fids []int) (map[int]string, error) {
	values := make([]string, 0, len(fids))
	for _, value := range fids {
		values = append(values, strconv.Itoa(value))
	}

	remarks := make(map[int]string)
	items, err := c.Rds.HMGet(ctx, c.name(uid), values...).Result()
	if err != nil {
		return nil, err
	}

	if len(items) == 0 {
		return remarks, nil
	}
	for k, v := range fids {
		if items[k] != nil {
			remarks[v] = items[k].(string)
		}
	}

	return remarks, nil
}

func (c *ContactRemarkCache) Set(ctx context.Context, uid int, friendId int, value string) error {
	if c.Exist(ctx, uid) {
		return c.Rds.HSet(ctx, c.name(uid), friendId, value).Err()
	}

	return nil
}

func (c *ContactRemarkCache) MSet(ctx context.Context, uid int, values map[string]any) error {
	_, err := c.Rds.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.HSet(ctx, c.name(uid), values)
		pipe.Expire(ctx, c.name(uid), 12*time.Hour)
		return nil
	})

	return err
}

func (c *ContactRemarkCache) Exist(ctx context.Context, uid int) bool {
	return c.Rds.Exists(ctx, c.name(uid)).Val() == 1
}

func (c *ContactRemarkCache) name(uid int) string {
	return fmt.Sprintf("im:contact:remark:uid_%d", uid)
}
