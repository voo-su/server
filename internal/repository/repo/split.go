package repo

import (
	"context"
	"gorm.io/gorm"
	"voo.su/internal/repository/model"
	"voo.su/pkg/core"
)

type Split struct {
	core.Repo[model.Split]
}

func NewFileSplit(db *gorm.DB) *Split {
	return &Split{Repo: core.NewRepo[model.Split](db)}
}

func (s *Split) GetSplitList(ctx context.Context, uploadId string) ([]*model.Split, error) {
	return s.Repo.FindAll(ctx, func(db *gorm.DB) {
		db.Where("upload_id = ? and type = 2", uploadId)
	})
}

func (s *Split) GetFile(ctx context.Context, uid int, uploadId string) (*model.Split, error) {
	return s.Repo.FindByWhere(ctx, "user_id = ? and upload_id = ? and type = 1", uid, uploadId)
}
