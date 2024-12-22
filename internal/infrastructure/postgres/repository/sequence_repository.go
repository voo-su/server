package repository

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
	"voo.su/internal/constant"
	postgresModel "voo.su/internal/infrastructure/postgres/model"
	redisRepo "voo.su/internal/infrastructure/redis/repository"
	"voo.su/pkg/logger"
	"voo.su/pkg/utils"
)

type SequenceRepository struct {
	DB                *gorm.DB
	SequenceCacheRepo *redisRepo.SequenceCacheRepository
}

func NewSequenceRepository(
	db *gorm.DB,
	sequenceCacheRepo *redisRepo.SequenceCacheRepository,
) *SequenceRepository {
	return &SequenceRepository{
		DB:                db,
		SequenceCacheRepo: sequenceCacheRepo,
	}
}

func (s *SequenceRepository) try(ctx context.Context, userId int, receiverId int) error {
	result := s.SequenceCacheRepo.Redis().TTL(ctx, s.SequenceCacheRepo.Name(userId, receiverId)).Val()
	if result == time.Duration(-2) {
		lockName := fmt.Sprintf("%s_lock", s.SequenceCacheRepo.Name(userId, receiverId))
		isTrue := s.SequenceCacheRepo.Redis().SetNX(ctx, lockName, 1, 10*time.Second).Val()
		if !isTrue {
			return errors.New("слишком частые запросы")
		}

		defer s.SequenceCacheRepo.Redis().Del(ctx, lockName)

		tx := s.DB.WithContext(ctx).Model(&postgresModel.Message{})
		if userId == 0 {
			tx = tx.Where("receiver_id = ? AND dialog_type = ?", receiverId, constant.ChatGroupMode)
		} else {
			tx = tx.Where("user_id = ? AND receiver_id = ?", userId, receiverId).
				Or("user_id = ? AND receiver_id = ?", receiverId, userId)
		}

		var seq int64
		if err := tx.
			Select("COALESCE(max(sequence),0)").
			Scan(&seq).
			Error; err != nil {
			logger.Errorf("Всего последовательностей, ошибка: %s", err.Error())
			return err
		}

		if err := s.SequenceCacheRepo.Set(ctx, userId, receiverId, seq); err != nil {
			logger.Errorf("Установка последовательности, ошибка: %s", err.Error())
			return err
		}
	} else if result < time.Hour {
		s.SequenceCacheRepo.Redis().Expire(ctx, s.SequenceCacheRepo.Name(userId, receiverId), 12*time.Hour)
	}

	return nil
}

func (s *SequenceRepository) Get(ctx context.Context, userId int, receiverId int) int64 {
	if err := utils.Retry(5, 100*time.Millisecond, func() error {
		return s.try(ctx, userId, receiverId)
	}); err != nil {
		log.Println("Ошибка получения последовательности: ", err.Error())
	}

	return s.SequenceCacheRepo.Get(ctx, userId, receiverId)
}

func (s *SequenceRepository) BatchGet(ctx context.Context, userId int, receiverId int, num int64) []int64 {
	if err := utils.Retry(5, 100*time.Millisecond, func() error {
		return s.try(ctx, userId, receiverId)
	}); err != nil {
		log.Println("Ошибка пакетного получения последовательности: ", err.Error())
	}

	return s.SequenceCacheRepo.BatchGet(ctx, userId, receiverId, num)
}
