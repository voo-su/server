package repository

import (
	"gorm.io/gorm"
	"voo.su/internal/infrastructure/postgres/model"
	"voo.su/pkg/gormutil"
)

type ContactFolderRepository struct {
	gormutil.Repo[model.ContactFolder]
}

func NewContactFolderRepository(db *gorm.DB) *ContactFolderRepository {
	return &ContactFolderRepository{Repo: gormutil.NewRepo[model.ContactFolder](db)}
}
