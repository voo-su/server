package repository

import (
	"context"
	"gorm.io/gorm"
	"voo.su/internal/infrastructure/postgres/model"
	"voo.su/pkg/repo"
)

type ChatRepository struct {
	repo.Repo[model.Chat]
}

func NewChatRepository(db *gorm.DB) *ChatRepository {
	return &ChatRepository{Repo: repo.NewRepo[model.Chat](db)}
}

func (c *ChatRepository) IsDisturb(uid int, receiverId int, chatType int) bool {
	resp, err := c.Repo.FindByWhere(context.TODO(), "user_id = ? AND receiver_id = ? AND dialog_type = ?", uid, receiverId, chatType)
	return err == nil && resp.IsDisturb == 1
}

func (c *ChatRepository) FindBySessionId(uid int, receiverId int, chatType int) int {
	resp, err := c.Repo.FindByWhere(context.TODO(), "user_id = ? AND receiver_id = ? AND dialog_type = ?", uid, receiverId, chatType)
	if err != nil {
		return 0
	}

	return resp.Id
}
