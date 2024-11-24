package repo

import (
	"context"
	"gorm.io/gorm"
	"voo.su/internal/repository/model"
	"voo.su/pkg/repo"
)

type Chat struct {
	repo.Repo[model.Chat]
}

func NewChat(db *gorm.DB) *Chat {
	return &Chat{Repo: repo.NewRepo[model.Chat](db)}
}

func (c *Chat) IsDisturb(uid int, receiverId int, chatType int) bool {
	resp, err := c.Repo.FindByWhere(context.TODO(), "user_id = ? AND receiver_id = ? AND dialog_type = ?", uid, receiverId, chatType)
	return err == nil && resp.IsDisturb == 1
}

func (c *Chat) FindBySessionId(uid int, receiverId int, chatType int) int {
	resp, err := c.Repo.FindByWhere(context.TODO(), "user_id = ? AND receiver_id = ? AND dialog_type = ?", uid, receiverId, chatType)
	if err != nil {
		return 0
	}

	return resp.Id
}
