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
	Conf        *config.Config
	ServerCache *cache.ServerCache
}

func NewHealthSubscribe(
	conf *config.Config,
	serverCache *cache.ServerCache,
) *HealthSubscribe {
	return &HealthSubscribe{
		Conf:        conf,
		ServerCache: serverCache,
	}
}

func (s *HealthSubscribe) Setup(ctx context.Context) error {
	log.Println("Запуск подписки на состояние здоровья")
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(10 * time.Second):
			if err := s.ServerCache.Set(ctx, s.Conf.ServerId(), time.Now().Unix()); err != nil {
				logger.Std().Error(fmt.Sprintf("Ошибка отчета о подписке на состояние WebSocket %s", err.Error()))
			}
		}
	}
}
