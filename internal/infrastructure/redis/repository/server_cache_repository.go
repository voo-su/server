package repository

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

type ServerCacheRepository struct {
	Rds *redis.Client
}

func NewServerCacheRepository(rds *redis.Client) *ServerCacheRepository {
	return &ServerCacheRepository{
		Rds: rds,
	}
}

func (s *ServerCacheRepository) Set(ctx context.Context, server string, time int64) error {
	_, err := s.Rds.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.SRem(ctx, ServerKeyExpire, server)
		pipe.HSet(ctx, ServerKey, server, time)
		return nil
	})
	return err
}

func (s *ServerCacheRepository) Del(ctx context.Context, server string) error {
	return s.Rds.HDel(ctx, ServerKey, server).Err()
}

func (s *ServerCacheRepository) All(ctx context.Context, status int) []string {
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

func (s *ServerCacheRepository) SetExpireServer(ctx context.Context, server string) error {
	return s.Rds.SAdd(ctx, ServerKeyExpire, server).Err()
}

func (s *ServerCacheRepository) DelExpireServer(ctx context.Context, server string) error {
	return s.Rds.SRem(ctx, ServerKeyExpire, server).Err()
}

func (s *ServerCacheRepository) GetExpireServerAll(ctx context.Context) []string {
	return s.Rds.SMembers(ctx, ServerKeyExpire).Val()
}

func (s *ServerCacheRepository) Redis() *redis.Client {
	return s.Rds
}
