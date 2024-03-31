package repo

import (
	"gorm.io/gorm"
	"voo.su/internal/repository/model"
	"voo.su/pkg/core"
)

type Message struct {
	core.Repo[model.Message]
}

func NewMessage(db *gorm.DB) *Message {
	return &Message{Repo: core.NewRepo[model.Message](db)}
}
