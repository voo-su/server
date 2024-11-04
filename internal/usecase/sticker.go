package usecase

import (
	"fmt"
	"strconv"
	"strings"
	"voo.su/internal/repository/model"
	"voo.su/internal/repository/repo"
	"voo.su/pkg/filesystem"
	"voo.su/pkg/sliceutil"
)

type StickerUseCase struct {
	*repo.Source
	StickerRepo *repo.Sticker
	Filesystem  *filesystem.Filesystem
}

func NewStickerUseCase(
	source *repo.Source,
	stickerRepo *repo.Sticker,
	fileSystem *filesystem.Filesystem,
) *StickerUseCase {
	return &StickerUseCase{
		Source:      source,
		StickerRepo: stickerRepo,
		Filesystem:  fileSystem,
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

	return s.Source.Db().Table("sticker_users").Where("user_id = ?", uid).
		Update("sticker_ids", strings.Join(items, ",")).
		Error
}

func (s *StickerUseCase) AddUserSysSticker(uid int, stickerId int) error {
	ids := s.StickerRepo.GetUserInstallIds(uid)
	if sliceutil.Include(stickerId, ids) {
		return nil
	}

	ids = append(ids, stickerId)
	return s.Source.Db().Table("sticker_users").
		Where("user_id = ?", uid).
		Update("sticker_ids", sliceutil.ToIds(ids)).
		Error
}

func (s *StickerUseCase) DeleteCollect(uid int, ids []int) error {
	return s.Source.Db().
		Delete(&model.StickerItem{}, "id in ? and sticker_id = 0 and user_id = ?", ids, uid).
		Error
}
