package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type Sequence struct {
	Rds *redis.Client
}

func NewSequence(rds *redis.Client) *Sequence {
	return &Sequence{Rds: rds}
}

func (s *Sequence) Redis() *redis.Client {
	return s.Rds
}

func (s *Sequence) Name(userId int, receiverId int) string {
	if userId == 0 {
		return fmt.Sprintf("im:sequence:chat:%d", receiverId)
	}

	if receiverId < userId {
		receiverId, userId = userId, receiverId
	}

	return fmt.Sprintf("im:sequence:chat:%d_%d", userId, receiverId)
}

func (s *Sequence) Set(ctx context.Context, userId int, receiverId int, value int64) error {
	return s.Rds.SetEx(ctx, s.Name(userId, receiverId), value, 12*time.Hour).Err()
}

func (s *Sequence) Get(ctx context.Context, userId int, receiverId int) int64 {
	return s.Rds.Incr(ctx, s.Name(userId, receiverId)).Val()
}

func (s *Sequence) BatchGet(ctx context.Context, userId int, receiverId int, num int64) []int64 {
	value := s.Rds.IncrBy(ctx, s.Name(userId, receiverId), num).Val()
	items := make([]int64, 0, num)
	for i := num; i > 0; i-- {
		items = append(items, value-i+1)
	}

	return items
}
