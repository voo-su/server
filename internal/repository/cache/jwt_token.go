package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
	"voo.su/pkg/encrypt"
)

type JwtTokenCache struct {
	Rds *redis.Client
}

func NewJwtTokenCache(redis *redis.Client) *JwtTokenCache {
	return &JwtTokenCache{redis}
}

func (j *JwtTokenCache) SetBlackList(ctx context.Context, token string, exp time.Duration) error {
	return j.Rds.Set(ctx, j.name(token), 1, exp).Err()
}

func (j *JwtTokenCache) IsBlackList(ctx context.Context, token string) bool {
	return j.Rds.Get(ctx, j.name(token)).Val() != ""
}

func (j *JwtTokenCache) name(token string) string {
	return fmt.Sprintf("jwt:blacklist:%s", encrypt.Md5(token))
}
