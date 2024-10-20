package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
	"voo.su/pkg/encrypt"
)

type SmsStorage struct {
	Redis *redis.Client
}

func NewSmsStorage(redis *redis.Client) *SmsStorage {
	return &SmsStorage{redis}
}

func (s *SmsStorage) Set(ctx context.Context, channel string, token string, code string, exp time.Duration) error {
	_, err := s.Redis.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.Del(ctx, s.failName(channel, token))
		pipe.Set(ctx, s.name(channel, token), code, exp)
		return nil
	})
	return err
}

func (s *SmsStorage) Get(ctx context.Context, channel string, token string) (string, error) {
	return s.Redis.Get(ctx, s.name(channel, token)).Result()
}

func (s *SmsStorage) Del(ctx context.Context, channel string, token string) error {
	return s.Redis.Del(ctx, s.name(channel, token)).Err()
}

func (s *SmsStorage) Verify(ctx context.Context, channel string, token string, code string) bool {
	value, err := s.Get(ctx, channel, token)
	if err != nil || len(value) == 0 {
		return false
	}
	if value == code {
		return true
	}

	num := s.Redis.Incr(ctx, s.failName(channel, token)).Val()
	if num >= 5 {
		_, _ = s.Redis.Pipelined(ctx, func(pipe redis.Pipeliner) error {
			pipe.Del(ctx, s.name(channel, token))
			pipe.Del(ctx, s.failName(channel, token))
			return nil
		})
	} else if num == 1 {
		s.Redis.Expire(ctx, s.failName(channel, token), 3*time.Minute)
	}

	return false
}

func (s *SmsStorage) name(channel string, token string) string {
	return fmt.Sprintf("im:auth:sms:%s:%s", channel, encrypt.Md5(token))
}

func (s *SmsStorage) failName(channel string, token string) string {
	return fmt.Sprintf("im:auth:sms_fail:%s:%s", channel, encrypt.Md5(token))
}
