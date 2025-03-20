package repository

import (
	"gorm.io/gorm"
	"voo.su/internal/infrastructure/postgres/model"
	"voo.su/pkg/gormutil"
)

type FileRepository struct {
	gormutil.Repo[model.File]
}

func NewFileRepository(db *gorm.DB) *FileRepository {
	return &FileRepository{Repo: gormutil.NewRepo[model.File](db)}
}
