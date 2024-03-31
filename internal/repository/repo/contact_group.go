package repo

import (
	"gorm.io/gorm"
	"voo.su/internal/repository/model"
	"voo.su/pkg/core"
)

type ContactGroup struct {
	core.Repo[model.ContactGroup]
}

func NewContactGroup(db *gorm.DB) *ContactGroup {
	return &ContactGroup{Repo: core.NewRepo[model.ContactGroup](db)}
}
