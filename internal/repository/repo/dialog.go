package repo

import (
	"context"
	"gorm.io/gorm"
	"voo.su/internal/repository/model"
	"voo.su/pkg/core"
)

type Dialog struct {
	core.Repo[model.Dialog]
}

func NewDialog(db *gorm.DB) *Dialog {
	return &Dialog{Repo: core.NewRepo[model.Dialog](db)}
}

func (d *Dialog) IsDisturb(uid int, receiverId int, dialogType int) bool {
	resp, err := d.Repo.FindByWhere(context.TODO(), "user_id = ? and receiver_id = ? and dialog_type = ?", uid, receiverId, dialogType)
	return err == nil && resp.IsDisturb == 1
}

func (d *Dialog) FindBySessionId(uid int, receiverId int, dialogType int) int {
	resp, err := d.Repo.FindByWhere(context.TODO(), "user_id = ? and receiver_id = ? and dialog_type = ?", uid, receiverId, dialogType)
	if err != nil {
		return 0
	}

	return resp.Id
}
