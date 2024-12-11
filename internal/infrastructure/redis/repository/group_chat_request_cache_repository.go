package repository

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type GroupChatRequestCacheRepository struct {
	Rds *redis.Client
}

func NewGroupChatRequestCacheRepository(rds *redis.Client) *GroupChatRequestCacheRepository {
	return &GroupChatRequestCacheRepository{rds}
}

func (g *GroupChatRequestCacheRepository) Incr(ctx context.Context, uid int) {
	g.Rds.Incr(ctx, g.name(uid))
}

func (g *GroupChatRequestCacheRepository) Get(ctx context.Context, uid int) int {
	val, err := g.Rds.Get(ctx, g.name(uid)).Int()
	if err != nil {
		return 0
	}

	return val
}

func (g *GroupChatRequestCacheRepository) Del(ctx context.Context, uid int) {
	g.Rds.Del(ctx, g.name(uid))
}

func (g *GroupChatRequestCacheRepository) name(uid int) string {
	return fmt.Sprintf("im:group:apply:unread:uid_%d", uid)
}
