package repository

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
	"voo.su/internal/config"
)

type ClientCacheRepository struct {
	Conf                  *config.Config
	Rds                   *redis.Client
	ServerCacheRepository *ServerCacheRepository
}

func NewClientCacheRepository(
	conf *config.Config,
	rds *redis.Client,
	serverCacheRepository *ServerCacheRepository,
) *ClientCacheRepository {
	return &ClientCacheRepository{
		Conf:                  conf,
		Rds:                   rds,
		ServerCacheRepository: serverCacheRepository,
	}
}

func (c *ClientCacheRepository) Set(ctx context.Context, channel string, fd string, uid int) error {
	sid := c.Conf.ServerId()
	_, err := c.Rds.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.HSet(ctx, c.clientKey(sid, channel), fd, uid)
		pipe.SAdd(ctx, c.userKey(sid, channel, strconv.Itoa(uid)), fd)
		return nil
	})

	return err
}

func (c *ClientCacheRepository) Del(ctx context.Context, channel, fd string) error {
	sid := c.Conf.ServerId()
	key := c.clientKey(sid, channel)
	uid, _ := c.Rds.HGet(ctx, key, fd).Result()
	_, err := c.Rds.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.HDel(ctx, key, fd)
		pipe.SRem(ctx, c.userKey(sid, channel, uid), fd)
		return nil
	})

	return err
}

func (c *ClientCacheRepository) IsOnline(ctx context.Context, channel, uid string) bool {
	for _, sid := range c.ServerCacheRepository.All(ctx, 1) {
		if c.IsCurrentServerOnline(ctx, sid, channel, uid) {
			return true
		}
	}
	return false
}

func (c *ClientCacheRepository) IsCurrentServerOnline(ctx context.Context, sid, channel, uid string) bool {
	val, err := c.Rds.SCard(ctx, c.userKey(sid, channel, uid)).Result()
	return err == nil && val > 0
}

func (c *ClientCacheRepository) GetUidFromClientIds(ctx context.Context, sid, channel, uid string) []int64 {
	cids := make([]int64, 0)
	items, err := c.Rds.SMembers(ctx, c.userKey(sid, channel, uid)).Result()
	if err != nil {
		return cids
	}

	for _, cid := range items {
		if cid, err := strconv.ParseInt(cid, 10, 64); err == nil {
			cids = append(cids, cid)
		}
	}

	return cids
}

func (c *ClientCacheRepository) GetClientIdFromUid(ctx context.Context, sid, channel, cid string) (int64, error) {
	uid, err := c.Rds.HGet(ctx, c.clientKey(sid, channel), cid).Result()
	if err != nil {
		return 0, err
	}

	return strconv.ParseInt(uid, 10, 64)
}

func (c *ClientCacheRepository) Bind(ctx context.Context, channel string, clientId int64, uid int) error {
	return c.Set(ctx, channel, strconv.FormatInt(clientId, 10), uid)
}

func (c *ClientCacheRepository) UnBind(ctx context.Context, channel string, clientId int64) error {
	return c.Del(ctx, channel, strconv.FormatInt(clientId, 10))
}

func (c *ClientCacheRepository) clientKey(sid, channel string) string {
	return fmt.Sprintf("ws:%s:%s:client", sid, channel)
}

func (c *ClientCacheRepository) userKey(sid, channel, uid string) string {
	return fmt.Sprintf("ws:%s:%s:user:%s", sid, channel, uid)
}
