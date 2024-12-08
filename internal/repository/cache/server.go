package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

const (
	ServerKey       = "server_ids"
	ServerKeyExpire = "server_ids_expire"
	ServerOverTime  = 50
)

type ServerCache struct {
	Rds *redis.Client
}

func NewServerCache(rds *redis.Client) *ServerCache {
	return &ServerCache{Rds: rds}
}

func (s *ServerCache) Set(ctx context.Context, server string, time int64) error {
	_, err := s.Rds.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.SRem(ctx, ServerKeyExpire, server)
		pipe.HSet(ctx, ServerKey, server, time)
		return nil
	})
	return err
}

func (s *ServerCache) Del(ctx context.Context, server string) error {
	return s.Rds.HDel(ctx, ServerKey, server).Err()
}

func (s *ServerCache) All(ctx context.Context, status int) []string {
	var (
		unix  = time.Now().Unix()
		slice = make([]string, 0)
	)
	all, err := s.Rds.HGetAll(ctx, ServerKey).Result()
	if err != nil {
		return slice
	}

	for key, val := range all {
		value, err := strconv.Atoi(val)
		if err != nil {
			continue
		}
		switch status {
		case 1:
			if unix-int64(value) >= ServerOverTime {
				continue
			}
		case 2:
			if unix-int64(value) < ServerOverTime {
				continue
			}
		}
		slice = append(slice, key)
	}
	return slice
}

func (s *ServerCache) SetExpireServer(ctx context.Context, server string) error {
	return s.Rds.SAdd(ctx, ServerKeyExpire, server).Err()
}

func (s *ServerCache) DelExpireServer(ctx context.Context, server string) error {
	return s.Rds.SRem(ctx, ServerKeyExpire, server).Err()
}

func (s *ServerCache) GetExpireServerAll(ctx context.Context) []string {
	return s.Rds.SMembers(ctx, ServerKeyExpire).Val()
}

func (s *ServerCache) Redis() *redis.Client {
	return s.Rds
}
