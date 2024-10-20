package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type GroupChatRequestStorage struct {
	Redis *redis.Client
}

func NewGroupChatRequestStorage(rds *redis.Client) *GroupChatRequestStorage {
	return &GroupChatRequestStorage{rds}
}

func (g *GroupChatRequestStorage) Incr(ctx context.Context, uid int) {
	g.Redis.Incr(ctx, g.name(uid))
}

func (g *GroupChatRequestStorage) Get(ctx context.Context, uid int) int {
	val, err := g.Redis.Get(ctx, g.name(uid)).Int()
	if err != nil {
		return 0
	}

	return val
}

func (g *GroupChatRequestStorage) Del(ctx context.Context, uid int) {
	g.Redis.Del(ctx, g.name(uid))
}

func (g *GroupChatRequestStorage) name(uid int) string {
	return fmt.Sprintf("im:group:apply:unread:uid_%d", uid)
}
