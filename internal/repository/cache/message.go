package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"voo.su/pkg/jsonutil"
)

const lastMessageCacheKey = "redis:hash:last-message"

type MessageStorage struct {
	Redis *redis.Client
}

type LastCacheMessage struct {
	Content  string `json:"content"`
	Datetime string `json:"datetime"`
}

func NewMessageStorage(rds *redis.Client) *MessageStorage {
	return &MessageStorage{rds}
}

func (m *MessageStorage) Set(ctx context.Context, dialogType int, sender int, receive int, message *LastCacheMessage) error {
	text := jsonutil.Encode(message)
	return m.Redis.HSet(ctx, lastMessageCacheKey, m.name(dialogType, sender, receive), text).Err()
}

func (m *MessageStorage) Get(ctx context.Context, dialogType int, sender int, receive int) (*LastCacheMessage, error) {
	res, err := m.Redis.HGet(ctx, lastMessageCacheKey, m.name(dialogType, sender, receive)).Result()
	if err != nil {
		return nil, err
	}

	msg := &LastCacheMessage{}
	if err = jsonutil.Decode(res, msg); err != nil {
		return nil, err
	}

	return msg, nil
}

func (m *MessageStorage) MGet(ctx context.Context, fields []string) ([]*LastCacheMessage, error) {
	res := m.Redis.HMGet(ctx, lastMessageCacheKey, fields...)
	items := make([]*LastCacheMessage, 0)
	for _, item := range res.Val() {
		if val, ok := item.(string); ok {
			msg := &LastCacheMessage{}
			if err := jsonutil.Decode(val, msg); err != nil {
				return nil, err
			}

			items = append(items, msg)
		}
	}

	return items, nil
}

func (m *MessageStorage) name(dialogType int, sender int, receive int) string {
	if dialogType == 2 {
		sender = 0
	}

	if sender > receive {
		sender, receive = receive, sender
	}

	return fmt.Sprintf("%d_%d_%d", dialogType, sender, receive)
}
