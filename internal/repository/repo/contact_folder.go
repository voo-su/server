package repo

import (
	"gorm.io/gorm"
	"voo.su/internal/repository/model"
	"voo.su/pkg/repo"
)

type ContactFolder struct {
	repo.Repo[model.ContactFolder]
}

func NewContactFolder(db *gorm.DB) *ContactFolder {
	return &ContactFolder{Repo: repo.NewRepo[model.ContactFolder](db)}
}
