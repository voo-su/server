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

type JwtTokenCacheRepository struct {
	Rds *redis.Client
}

func NewJwtTokenCacheRepository(rds *redis.Client) *JwtTokenCacheRepository {
	return &JwtTokenCacheRepository{
		Rds: rds,
	}
}

func (j *JwtTokenCacheRepository) SetBlackList(ctx context.Context, token string, exp time.Duration) error {
	return j.Rds.Set(ctx, j.name(token), 1, exp).Err()
}

func (j *JwtTokenCacheRepository) IsBlackList(ctx context.Context, token string) bool {
	return j.Rds.Get(ctx, j.name(token)).Val() != ""
}

func (j *JwtTokenCacheRepository) name(token string) string {
	return fmt.Sprintf("jwt:blacklist:%s", encrypt.Md5(token))
}
