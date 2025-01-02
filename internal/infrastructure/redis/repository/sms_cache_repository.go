// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package repository

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
	"voo.su/pkg/encrypt"
)

type SmsCacheRepository struct {
	Rds *redis.Client
}

func NewSmsCacheRepository(rds *redis.Client) *SmsCacheRepository {
	return &SmsCacheRepository{
		Rds: rds,
	}
}

func (s *SmsCacheRepository) Set(ctx context.Context, channel string, token string, code string, exp time.Duration) error {
	_, err := s.Rds.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.Del(ctx, s.failName(channel, token))
		pipe.Set(ctx, s.name(channel, token), code, exp)
		return nil
	})
	return err
}

func (s *SmsCacheRepository) Get(ctx context.Context, channel string, token string) (string, error) {
	return s.Rds.Get(ctx, s.name(channel, token)).Result()
}

func (s *SmsCacheRepository) Del(ctx context.Context, channel string, token string) error {
	return s.Rds.Del(ctx, s.name(channel, token)).Err()
}

func (s *SmsCacheRepository) Verify(ctx context.Context, channel string, token string, code string) bool {
	value, err := s.Get(ctx, channel, token)
	if err != nil || len(value) == 0 {
		return false
	}
	if value == code {
		return true
	}

	num := s.Rds.Incr(ctx, s.failName(channel, token)).Val()
	if num >= 5 {
		_, _ = s.Rds.Pipelined(ctx, func(pipe redis.Pipeliner) error {
			pipe.Del(ctx, s.name(channel, token))
			pipe.Del(ctx, s.failName(channel, token))
			return nil
		})
	} else if num == 1 {
		s.Rds.Expire(ctx, s.failName(channel, token), 3*time.Minute)
	}

	return false
}

func (s *SmsCacheRepository) name(channel string, token string) string {
	return fmt.Sprintf("im:auth:sms:%s:%s", channel, encrypt.Md5(token))
}

func (s *SmsCacheRepository) failName(channel string, token string) string {
	return fmt.Sprintf("im:auth:sms_fail:%s:%s", channel, encrypt.Md5(token))
}
