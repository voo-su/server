package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
	"voo.su/internal/entity"
)

type RoomStorage struct {
	redis *redis.Client
}

type RoomOption struct {
	Channel  string
	RoomType entity.RoomType
	Number   string
	Sid      string
	Cid      int64
}

func NewRoomStorage(redis *redis.Client) *RoomStorage {
	return &RoomStorage{redis: redis}
}

func (r *RoomStorage) Add(ctx context.Context, opt *RoomOption) error {
	key := r.name(opt)
	err := r.redis.SAdd(ctx, key, opt.Cid).Err()
	if err == nil {
		r.redis.Expire(ctx, key, time.Hour*24*7)
	}

	return err
}

func (r *RoomStorage) BatchAdd(ctx context.Context, opts []*RoomOption) error {
	pipeline := r.redis.Pipeline()
	for _, opt := range opts {
		key := r.name(opt)
		if err := pipeline.SAdd(ctx, key, opt.Cid).Err(); err == nil {
			pipeline.Expire(ctx, key, time.Hour*24*7)
		}
	}
	_, err := pipeline.Exec(ctx)

	return err
}

func (r *RoomStorage) Del(ctx context.Context, opt *RoomOption) error {
	return r.redis.SRem(ctx, r.name(opt), opt.Cid).Err()
}

func (r *RoomStorage) BatchDel(ctx context.Context, opts []*RoomOption) error {
	pipeline := r.redis.Pipeline()
	for _, opt := range opts {
		pipeline.SRem(ctx, r.name(opt), opt.Cid)
	}

	_, err := pipeline.Exec(ctx)
	return err
}

func (r *RoomStorage) All(ctx context.Context, opt *RoomOption) []int64 {
	arr := r.redis.SMembers(ctx, r.name(opt)).Val()
	cids := make([]int64, 0, len(arr))
	for _, val := range arr {
		if cid, err := strconv.ParseInt(val, 10, 64); err == nil {
			cids = append(cids, cid)
		}
	}

	return cids
}

func (r *RoomStorage) name(opt *RoomOption) string {
	return fmt.Sprintf("ws:%s:%s:%s", opt.Sid, opt.RoomType, opt.Number)
}
