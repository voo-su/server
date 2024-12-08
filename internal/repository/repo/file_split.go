package repo

import (
	"context"
	"gorm.io/gorm"
	"voo.su/internal/repository/model"
	"voo.su/pkg/repo"
)

type FileSplit struct {
	repo.Repo[model.FileSplit]
}

func NewFileSplit(db *gorm.DB) *FileSplit {
	return &FileSplit{Repo: repo.NewRepo[model.FileSplit](db)}
}

func (s *FileSplit) GetSplitList(ctx context.Context, uploadId string) ([]*model.FileSplit, error) {
	return s.Repo.FindAll(ctx, func(db *gorm.DB) {
		db.Where("upload_id = ? AND type = 2", uploadId)
	})
}

func (s *FileSplit) GetFile(ctx context.Context, uid int, uploadId string) (*model.FileSplit, error) {
	return s.Repo.FindByWhere(ctx, "user_id = ? AND upload_id = ? AND type = 1", uid, uploadId)
}
