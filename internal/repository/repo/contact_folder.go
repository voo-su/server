package repo

import (
	"gorm.io/gorm"
	"voo.su/internal/repository/model"
	"voo.su/pkg/core"
)

type ContactFolder struct {
	core.Repo[model.ContactFolder]
}

func NewContactFolder(db *gorm.DB) *ContactFolder {
	return &ContactFolder{Repo: core.NewRepo[model.ContactFolder](db)}
}
