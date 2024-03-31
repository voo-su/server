package process

import (
	"context"
	"fmt"
	"log"
	"time"
	"voo.su/internal/config"
	"voo.su/internal/repository/cache"
	"voo.su/pkg/logger"
)

type HealthSubscribe struct {
	config  *config.Config
	storage *cache.ServerStorage
}

func NewHealthSubscribe(config *config.Config, storage *cache.ServerStorage) *HealthSubscribe {
	return &HealthSubscribe{config: config, storage: storage}
}

func (s *HealthSubscribe) Setup(ctx context.Context) error {
	log.Println("Запуск подписки на состояние здоровья")
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(10 * time.Second):
			if err := s.storage.Set(ctx, s.config.ServerId(), time.Now().Unix()); err != nil {
				logger.Std().Error(fmt.Sprintf("Ошибка отчета о подписке на состояние WebSocket %s", err.Error()))
			}
		}
	}
}
