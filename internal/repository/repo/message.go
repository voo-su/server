package repo

import (
	"gorm.io/gorm"
	"voo.su/internal/repository/model"
	"voo.su/pkg/repo"
)

type Message struct {
	repo.Repo[model.Message]
}

func NewMessage(db *gorm.DB) *Message {
	return &Message{Repo: repo.NewRepo[model.Message](db)}
}
