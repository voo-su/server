package repository

import (
	"gorm.io/gorm"
	"voo.su/internal/infrastructure/postgres/model"
	"voo.su/pkg/repo"
)

type ContactFolderRepository struct {
	repo.Repo[model.ContactFolder]
}

func NewContactFolderRepository(db *gorm.DB) *ContactFolderRepository {
	return &ContactFolderRepository{Repo: repo.NewRepo[model.ContactFolder](db)}
}
