package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
	"voo.su/pkg/encrypt"
)

type JwtTokenStorage struct {
	Redis *redis.Client
}

func NewTokenSessionStorage(redis *redis.Client) *JwtTokenStorage {
	return &JwtTokenStorage{redis}
}

func (s *JwtTokenStorage) SetBlackList(ctx context.Context, token string, exp time.Duration) error {
	return s.Redis.Set(ctx, s.name(token), 1, exp).Err()
}

func (s *JwtTokenStorage) IsBlackList(ctx context.Context, token string) bool {
	return s.Redis.Get(ctx, s.name(token)).Val() != ""
}

func (s *JwtTokenStorage) name(token string) string {
	return fmt.Sprintf("jwt:blacklist:%s", encrypt.Md5(token))
}
