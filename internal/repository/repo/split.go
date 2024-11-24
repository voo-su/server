package repo

import (
	"context"
	"gorm.io/gorm"
	"voo.su/internal/repository/model"
	"voo.su/pkg/repo"
)

type Split struct {
	repo.Repo[model.Split]
}

func NewFileSplit(db *gorm.DB) *Split {
	return &Split{Repo: repo.NewRepo[model.Split](db)}
}

func (s *Split) GetSplitList(ctx context.Context, uploadId string) ([]*model.Split, error) {
	return s.Repo.FindAll(ctx, func(db *gorm.DB) {
		db.Where("upload_id = ? AND type = 2", uploadId)
	})
}

func (s *Split) GetFile(ctx context.Context, uid int, uploadId string) (*model.Split, error) {
	return s.Repo.FindByWhere(ctx, "user_id = ? AND upload_id = ? AND type = 1", uid, uploadId)
}
