// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package process

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/sourcegraph/conc/pool"
	"log"
	"voo.su/internal/config"
	"voo.su/internal/constant"
	"voo.su/internal/delivery/ws/consume"
	"voo.su/pkg/utils"
)

type MessageSubscribe struct {
	Conf           *config.Config
	Redis          *redis.Client
	DefaultConsume *consume.ChatSubscribe
}

func NewMessageSubscribe(
	conf *config.Config,
	redis *redis.Client,
	defaultConsume *consume.ChatSubscribe,
) *MessageSubscribe {
	return &MessageSubscribe{
		Conf:           conf,
		Redis:          redis,
		DefaultConsume: defaultConsume,
	}
}

type IConsume interface {
	Call(event string, data []byte)
}

func (m *MessageSubscribe) Setup(ctx context.Context) error {
	go m.subscribe(ctx, []string{
		constant.ImTopicChat,
		fmt.Sprintf(constant.ImTopicChatPrivate,
			m.Conf.ServerId())}, m.DefaultConsume)
	<-ctx.Done()
	return nil
}

type SubscribeContent struct {
	Event string `json:"event"`
	Data  string `json:"data"`
}

func (m *MessageSubscribe) subscribe(ctx context.Context, topic []string, consume IConsume) {
	sub := m.Redis.Subscribe(ctx, topic...)
	defer sub.Close()
	worker := pool.New().WithMaxGoroutines(10)
	for data := range sub.Channel() {
		m.handle(worker, data, consume)
	}
	worker.Wait()
}

func (m *MessageSubscribe) handle(worker *pool.Pool, data *redis.Message, consume IConsume) {
	worker.Go(func() {
		var in SubscribeContent
		if err := json.Unmarshal([]byte(data.Payload), &in); err != nil {
			log.Printf("Content unsubscription error: %s", err)
			return
		}
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Error subscribing to call notification: %s", utils.PanicTrace(err))
			}
		}()
		consume.Call(in.Event, []byte(in.Data))
	})
}
