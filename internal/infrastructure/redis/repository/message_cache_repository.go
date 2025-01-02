// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package repository

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"voo.su/internal/infrastructure/redis/model"
	"voo.su/pkg/jsonutil"
)

const lastMessageCacheKey = "redis:hash:last-message"

type MessageCacheRepository struct {
	Rds *redis.Client
}

func NewMessageCacheRepository(rds *redis.Client) *MessageCacheRepository {
	return &MessageCacheRepository{
		Rds: rds,
	}
}

func (m *MessageCacheRepository) Set(ctx context.Context, dialogType int, sender int, receive int, message *model.LastCacheMessage) error {
	text := jsonutil.Encode(message)
	return m.Rds.HSet(ctx, lastMessageCacheKey, m.name(dialogType, sender, receive), text).Err()
}

func (m *MessageCacheRepository) Get(ctx context.Context, dialogType int, sender int, receive int) (*model.LastCacheMessage, error) {
	res, err := m.Rds.HGet(ctx, lastMessageCacheKey, m.name(dialogType, sender, receive)).Result()
	if err != nil {
		return nil, err
	}

	msg := &model.LastCacheMessage{}
	if err = jsonutil.Decode(res, msg); err != nil {
		return nil, err
	}

	return msg, nil
}

func (m *MessageCacheRepository) MGet(ctx context.Context, fields []string) ([]*model.LastCacheMessage, error) {
	res := m.Rds.HMGet(ctx, lastMessageCacheKey, fields...)
	items := make([]*model.LastCacheMessage, 0)
	for _, item := range res.Val() {
		if val, ok := item.(string); ok {
			msg := &model.LastCacheMessage{}
			if err := jsonutil.Decode(val, msg); err != nil {
				return nil, err
			}

			items = append(items, msg)
		}
	}

	return items, nil
}

func (m *MessageCacheRepository) name(dialogType int, sender int, receive int) string {
	if dialogType == 2 {
		sender = 0
	}

	if sender > receive {
		sender, receive = receive, sender
	}

	return fmt.Sprintf("%d_%d_%d", dialogType, sender, receive)
}
