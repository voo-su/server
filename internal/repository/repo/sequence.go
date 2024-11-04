package repo

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
	"voo.su/internal/constant"
	"voo.su/internal/repository/cache"
	"voo.su/internal/repository/model"
	"voo.su/pkg/logger"
	"voo.su/pkg/utils"
)

type Sequence struct {
	db    *gorm.DB
	cache *cache.Sequence
}

func NewSequence(db *gorm.DB, cache *cache.Sequence) *Sequence {
	return &Sequence{db: db, cache: cache}
}

func (s *Sequence) try(ctx context.Context, userId int, receiverId int) error {
	result := s.cache.Redis().TTL(ctx, s.cache.Name(userId, receiverId)).Val()
	if result == time.Duration(-2) {
		lockName := fmt.Sprintf("%s_lock", s.cache.Name(userId, receiverId))
		isTrue := s.cache.Redis().SetNX(ctx, lockName, 1, 10*time.Second).Val()
		if !isTrue {
			return errors.New("слишком частые запросы")
		}

		defer s.cache.Redis().Del(ctx, lockName)
		tx := s.db.WithContext(ctx).Model(&model.Message{})
		if userId == 0 {
			tx = tx.Where("receiver_id = ? and dialog_type = ?", receiverId, constant.ChatGroupMode)
		} else {
			tx = tx.Where("user_id = ? and receiver_id = ?", userId, receiverId).Or("user_id = ? and receiver_id = ?", receiverId, userId)
		}

		var seq int64
		err := tx.Select("COALESCE(max(sequence),0)").Scan(&seq).Error
		if err != nil {
			logger.Errorf("Всего последовательностей, ошибка: %s", err.Error())
			return err
		}

		if err := s.cache.Set(ctx, userId, receiverId, seq); err != nil {
			logger.Errorf("Установка последовательности, ошибка: %s", err.Error())
			return err
		}
	} else if result < time.Hour {
		s.cache.Redis().Expire(ctx, s.cache.Name(userId, receiverId), 12*time.Hour)
	}

	return nil
}

func (s *Sequence) Get(ctx context.Context, userId int, receiverId int) int64 {
	if err := utils.Retry(5, 100*time.Millisecond, func() error {
		return s.try(ctx, userId, receiverId)
	}); err != nil {
		log.Println("Ошибка получения последовательности: ", err.Error())
	}

	return s.cache.Get(ctx, userId, receiverId)
}

func (s *Sequence) BatchGet(ctx context.Context, userId int, receiverId int, num int64) []int64 {
	if err := utils.Retry(5, 100*time.Millisecond, func() error {
		return s.try(ctx, userId, receiverId)
	}); err != nil {
		log.Println("Ошибка пакетного получения последовательности: ", err.Error())
	}

	return s.cache.BatchGet(ctx, userId, receiverId, num)
}
