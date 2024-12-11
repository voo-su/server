package repository

import (
	"context"
	"gorm.io/gorm"
	model2 "voo.su/internal/infrastructure/postgres/model"
	"voo.su/pkg/repo"
	"voo.su/pkg/sliceutil"
)

type StickerRepository struct {
	repo.Repo[model2.Sticker]
}

func NewStickerRepository(db *gorm.DB) *StickerRepository {
	return &StickerRepository{Repo: repo.NewRepo[model2.Sticker](db)}
}

func (e *StickerRepository) GetUserInstallIds(uid int) []int {
	var data model2.StickerUser
	if err := e.Repo.Db.First(&data, "user_id = ?", uid).Error; err != nil {
		return []int{}
	}

	return sliceutil.ParseIds(data.StickerIds)
}

func (e *StickerRepository) GetSystemStickerList(ctx context.Context) ([]*model2.Sticker, error) {
	return e.Repo.FindAll(ctx, func(db *gorm.DB) {
		db.Where("status = ?", 0)
	})
}

func (e *StickerRepository) GetDetailsAll(stickerId, uid int) ([]*model2.StickerItem, error) {
	var items []*model2.StickerItem
	if err := e.Repo.Db.Model(model2.StickerItem{}).
		Where("sticker_id = ? AND user_id = ? order by id desc", stickerId, uid).
		Scan(&items).
		Error; err != nil {
		return nil, err
	}

	return items, nil
}
