package process

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/sourcegraph/conc/pool"
	"log"
	"voo.su/internal/config"
	"voo.su/internal/entity"
	"voo.su/internal/transport/ws/consume"
	"voo.su/pkg/utils"
)

type MessageSubscribe struct {
	config         *config.Config
	redis          *redis.Client
	defaultConsume *consume.ChatSubscribe
}

func NewMessageSubscribe(config *config.Config, redis *redis.Client, defaultConsume *consume.ChatSubscribe) *MessageSubscribe {
	return &MessageSubscribe{config: config, redis: redis, defaultConsume: defaultConsume}
}

type IConsume interface {
	Call(event string, data []byte)
}

func (m *MessageSubscribe) Setup(ctx context.Context) error {
	log.Println("Старт подписки на сообщение")
	go m.subscribe(ctx, []string{entity.ImTopicChat, fmt.Sprintf(entity.ImTopicChatPrivate, m.config.ServerId())}, m.defaultConsume)
	<-ctx.Done()
	return nil
}

type SubscribeContent struct {
	Event string `json:"event"`
	Data  string `json:"data"`
}

func (m *MessageSubscribe) subscribe(ctx context.Context, topic []string, consume IConsume) {
	sub := m.redis.Subscribe(ctx, topic...)
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
			log.Println("Ошибка отмены подписки на контент: ", err.Error())
			return
		}
		defer func() {
			if err := recover(); err != nil {
				log.Println("Ошибка при подписке на сообщение о вызове: ", utils.PanicTrace(err))
			}
		}()
		consume.Call(in.Event, []byte(in.Data))
	})
}
