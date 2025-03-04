package repository

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
	"voo.su/internal/infrastructure/redis/model"
)

type RoomCacheRepository struct {
	Rds *redis.Client
}

func NewRoomCacheRepository(rds *redis.Client) *RoomCacheRepository {
	return &RoomCacheRepository{
		Rds: rds,
	}
}

func (r *RoomCacheRepository) Add(ctx context.Context, opt *model.RoomOption) error {
	key := r.name(opt)
	err := r.Rds.SAdd(ctx, key, opt.Cid).Err()
	if err == nil {
		r.Rds.Expire(ctx, key, time.Hour*24*7)
	}

	return err
}

func (r *RoomCacheRepository) BatchAdd(ctx context.Context, opts []*model.RoomOption) error {
	pipeline := r.Rds.Pipeline()
	for _, opt := range opts {
		key := r.name(opt)
		if err := pipeline.SAdd(ctx, key, opt.Cid).Err(); err == nil {
			pipeline.Expire(ctx, key, time.Hour*24*7)
		}
	}
	_, err := pipeline.Exec(ctx)

	return err
}

func (r *RoomCacheRepository) Del(ctx context.Context, opt *model.RoomOption) error {
	return r.Rds.SRem(ctx, r.name(opt), opt.Cid).Err()
}

func (r *RoomCacheRepository) BatchDel(ctx context.Context, opts []*model.RoomOption) error {
	pipeline := r.Rds.Pipeline()
	for _, opt := range opts {
		pipeline.SRem(ctx, r.name(opt), opt.Cid)
	}

	_, err := pipeline.Exec(ctx)
	return err
}

func (r *RoomCacheRepository) All(ctx context.Context, opt *model.RoomOption) []int64 {
	arr := r.Rds.SMembers(ctx, r.name(opt)).Val()
	cids := make([]int64, 0, len(arr))
	for _, val := range arr {
		if cid, err := strconv.ParseInt(val, 10, 64); err == nil {
			cids = append(cids, cid)
		}
	}

	return cids
}

func (r *RoomCacheRepository) name(opt *model.RoomOption) string {
	return fmt.Sprintf("ws:%s:%s:%s", opt.Sid, opt.RoomType, opt.Number)
}
