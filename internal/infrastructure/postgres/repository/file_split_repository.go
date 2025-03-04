package repository

import (
	"context"
	"gorm.io/gorm"
	"voo.su/internal/infrastructure/postgres/model"
	"voo.su/pkg/gormutil"
)

type FileSplitRepository struct {
	gormutil.Repo[model.FileSplit]
}

func NewFileSplitRepository(db *gorm.DB) *FileSplitRepository {
	return &FileSplitRepository{Repo: gormutil.NewRepo[model.FileSplit](db)}
}

func (s *FileSplitRepository) GetSplitList(ctx context.Context, uploadId string) ([]*model.FileSplit, error) {
	return s.Repo.FindAll(ctx, func(db *gorm.DB) {
		db.Where("upload_id = ? AND type = ?", uploadId, 2)
	})
}

func (s *FileSplitRepository) GetFile(ctx context.Context, uid int, uploadId string) (*model.FileSplit, error) {
	return s.Repo.FindByWhere(ctx, "user_id = ? AND upload_id = ? AND type = ?", uid, uploadId, 1)
}
