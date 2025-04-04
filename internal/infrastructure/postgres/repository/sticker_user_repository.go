package repository

import (
	"context"
	"gorm.io/gorm"
	"voo.su/internal/infrastructure/postgres/model"
	"voo.su/pkg/gormutil"
	"voo.su/pkg/sliceutil"
)

type StickerRepository struct {
	gormutil.Repo[model.Sticker]
}

func NewStickerRepository(db *gorm.DB) *StickerRepository {
	return &StickerRepository{Repo: gormutil.NewRepo[model.Sticker](db)}
}

func (e *StickerRepository) GetUserInstallIds(uid int) []int {
	var data model.StickerUser
	if err := e.Repo.Db.First(&data, "user_id = ?", uid).Error; err != nil {
		return []int{}
	}

	return sliceutil.ParseIds(data.StickerIds)
}

func (e *StickerRepository) GetSystemStickerList(ctx context.Context) ([]*model.Sticker, error) {
	return e.Repo.FindAll(ctx, func(db *gorm.DB) {
		db.Where("status = ?", 0)
	})
}

func (e *StickerRepository) GetDetailsAll(stickerId, uid int) ([]*model.StickerItem, error) {
	var items []*model.StickerItem
	if err := e.Repo.Db.Model(model.StickerItem{}).
		Where("sticker_id = ? AND user_id = ? order by id desc", stickerId, uid).
		Scan(&items).
		Error; err != nil {
		return nil, err
	}

	return items, nil
}
