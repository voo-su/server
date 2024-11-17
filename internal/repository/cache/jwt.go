package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
	"voo.su/pkg/encrypt"
)

type JwtTokenStorage struct {
	Rds *redis.Client
}

func NewTokenSessionStorage(redis *redis.Client) *JwtTokenStorage {
	return &JwtTokenStorage{redis}
}

func (j *JwtTokenStorage) SetBlackList(ctx context.Context, token string, exp time.Duration) error {
	return j.Rds.Set(ctx, j.name(token), 1, exp).Err()
}

func (j *JwtTokenStorage) IsBlackList(ctx context.Context, token string) bool {
	return j.Rds.Get(ctx, j.name(token)).Val() != ""
}

func (j *JwtTokenStorage) name(token string) string {
	return fmt.Sprintf("jwt:blacklist:%s", encrypt.Md5(token))
}
