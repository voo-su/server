package usecase

import (
	"fmt"
	"strconv"
	"strings"
	"voo.su/internal/infrastructure"
	"voo.su/internal/infrastructure/postgres/model"
	postgresRepo "voo.su/internal/infrastructure/postgres/repository"
	redisRepo "voo.su/internal/infrastructure/redis/repository"
	"voo.su/pkg/locale"
	"voo.su/pkg/minio"
	"voo.su/pkg/sliceutil"
)

type StickerUseCase struct {
	Locale             locale.ILocale
	Source             *infrastructure.Source
	StickerRepo        *postgresRepo.StickerRepository
	IMinio             minio.IMinio
	RedisLockCacheRepo *redisRepo.RedisLockCacheRepository
}

func NewStickerUseCase(
	locale locale.ILocale,
	source *infrastructure.Source,
	stickerRepo *postgresRepo.StickerRepository,
	minio minio.IMinio,
	redisLockCacheRepo *redisRepo.RedisLockCacheRepository,
) *StickerUseCase {
	return &StickerUseCase{
		Locale:             locale,
		Source:             source,
		StickerRepo:        stickerRepo,
		IMinio:             minio,
		RedisLockCacheRepo: redisLockCacheRepo,
	}
}

func (s *StickerUseCase) RemoveUserSysSticker(uid int, stickerId int) error {
	ids := s.StickerRepo.GetUserInstallIds(uid)
	if !sliceutil.Include(stickerId, ids) {
		return fmt.Errorf(s.Locale.Localize("data_not_found"))
	}

	items := make([]string, 0, len(ids)-1)
	for _, id := range ids {
		if id != stickerId {
			items = append(items, strconv.Itoa(id))
		}
	}

	return s.Source.Postgres().
		Table("sticker_users").
		Where("user_id = ?", uid).
		Update("sticker_ids", strings.Join(items, ",")).
		Error
}

func (s *StickerUseCase) AddUserSysSticker(uid int, stickerId int) error {
	ids := s.StickerRepo.GetUserInstallIds(uid)
	if sliceutil.Include(stickerId, ids) {
		return nil
	}

	ids = append(ids, stickerId)
	return s.Source.Postgres().
		Table("sticker_users").
		Where("user_id = ?", uid).
		Update("sticker_ids", sliceutil.ToIds(ids)).
		Error
}

func (s *StickerUseCase) DeleteCollect(uid int, ids []int) error {
	return s.Source.Postgres().
		Delete(&model.StickerItem{}, "id IN ? AND sticker_id = ? AND user_id = ?", ids, 0, uid).
		Error
}
