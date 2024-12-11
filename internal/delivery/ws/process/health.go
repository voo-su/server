package process

import (
	"context"
	"fmt"
	"log"
	"time"
	"voo.su/internal/config"
	redisRepo "voo.su/internal/infrastructure/redis/repository"
	"voo.su/pkg/logger"
)

type HealthSubscribe struct {
	Conf            *config.Config
	ServerCacheRepo *redisRepo.ServerCacheRepository
}

func NewHealthSubscribe(
	conf *config.Config,
	serverCacheRepo *redisRepo.ServerCacheRepository,
) *HealthSubscribe {
	return &HealthSubscribe{
		Conf:            conf,
		ServerCacheRepo: serverCacheRepo,
	}
}

func (s *HealthSubscribe) Setup(ctx context.Context) error {
	log.Println("Запуск подписки на состояние здоровья")
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(10 * time.Second):
			if err := s.ServerCacheRepo.Set(ctx, s.Conf.ServerId(), time.Now().Unix()); err != nil {
				logger.Std().Error(fmt.Sprintf("Ошибка отчета о подписке на состояние WebSocket %s", err.Error()))
			}
		}
	}
}
