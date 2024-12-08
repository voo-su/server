package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type GroupChatRequestCache struct {
	Rds *redis.Client
}

func NewGroupChatRequestCache(rds *redis.Client) *GroupChatRequestCache {
	return &GroupChatRequestCache{rds}
}

func (g *GroupChatRequestCache) Incr(ctx context.Context, uid int) {
	g.Rds.Incr(ctx, g.name(uid))
}

func (g *GroupChatRequestCache) Get(ctx context.Context, uid int) int {
	val, err := g.Rds.Get(ctx, g.name(uid)).Int()
	if err != nil {
		return 0
	}

	return val
}

func (g *GroupChatRequestCache) Del(ctx context.Context, uid int) {
	g.Rds.Del(ctx, g.name(uid))
}

func (g *GroupChatRequestCache) name(uid int) string {
	return fmt.Sprintf("im:group:apply:unread:uid_%d", uid)
}
