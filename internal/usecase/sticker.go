package usecase

import (
	"fmt"
	"strconv"
	"strings"
	"voo.su/internal/repository"
	"voo.su/internal/repository/model"
	"voo.su/internal/repository/repo"
	"voo.su/pkg/minio"
	"voo.su/pkg/sliceutil"
)

type StickerUseCase struct {
	*repository.Source
	StickerRepo *repo.Sticker
	IMinio      minio.IMinio
}

func NewStickerUseCase(
	source *repository.Source,
	stickerRepo *repo.Sticker,
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
		Delete(&model.StickerItem{}, "id in ? AND sticker_id = 0 AND user_id = ?", ids, uid).
		Error
}
