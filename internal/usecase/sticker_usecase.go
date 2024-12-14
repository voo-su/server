package usecase

import (
	"fmt"
	"strconv"
	"strings"
	"voo.su/internal/infrastructure"
	"voo.su/internal/infrastructure/postgres/model"
	postgresRepo "voo.su/internal/infrastructure/postgres/repository"
	"voo.su/pkg/minio"
	"voo.su/pkg/sliceutil"
)

type StickerUseCase struct {
	*infrastructure.Source
	StickerRepo *postgresRepo.StickerRepository
	IMinio      minio.IMinio
}

func NewStickerUseCase(
	source *infrastructure.Source,
	stickerRepo *postgresRepo.StickerRepository,
	minio minio.IMinio,
) *StickerUseCase {
	return &StickerUseCase{
		Source:      source,
		StickerRepo: stickerRepo,
		IMinio:      minio,
	}
}

func (s *StickerUseCase) RemoveUserSysSticker(uid int, stickerId int) error {
	ids := s.StickerRepo.GetUserInstallIds(uid)
	if !sliceutil.Include(stickerId, ids) {
		return fmt.Errorf("данных не существует")
	}

	items := make([]string, 0, len(ids)-1)
	for _, id := range ids {
		if id != stickerId {
			items = append(items, strconv.Itoa(id))
		}
	}

	return s.Source.Db().
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
	return s.Source.Db().
		Table("sticker_users").
		Where("user_id = ?", uid).
		Update("sticker_ids", sliceutil.ToIds(ids)).
		Error
}

func (s *StickerUseCase) DeleteCollect(uid int, ids []int) error {
	return s.Source.Db().
		Delete(&model.StickerItem{}, "id IN ? AND sticker_id = ? AND user_id = ?", ids, 0, uid).
		Error
}
