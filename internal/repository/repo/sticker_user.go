package repo

import (
	"context"
	"gorm.io/gorm"
	"voo.su/internal/repository/model"
	"voo.su/pkg/core"
	"voo.su/pkg/sliceutil"
)

type Sticker struct {
	core.Repo[model.Sticker]
}

func NewSticker(db *gorm.DB) *Sticker {
	return &Sticker{Repo: core.NewRepo[model.Sticker](db)}
}

func (e *Sticker) GetUserInstallIds(uid int) []int {
	var data model.MessageSticker
	if err := e.Repo.Db.First(&data, "user_id = ?", uid).Error; err != nil {
		return []int{}
	}

	return sliceutil.ParseIds(data.StickerIds)
}

func (e *Sticker) GetSystemStickerList(ctx context.Context) ([]*model.Sticker, error) {
	return e.Repo.FindAll(ctx, func(db *gorm.DB) {
		db.Where("status = ?", 0)
	})
}

func (e *Sticker) GetDetailsAll(stickerId, uid int) ([]*model.StickerItem, error) {
	var items []*model.StickerItem
	if err := e.Repo.Db.Model(model.StickerItem{}).
		Where("sticker_id = ? and user_id = ? order by id desc", stickerId, uid).
		Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}
