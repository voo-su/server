package model

import "time"

type Chat struct {
	Id         int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	DialogType int       `gorm:"column:dialog_type;default:1;NOT NULL" json:"dialog_type"`
	UserId     int       `gorm:"column:user_id;default:0;NOT NULL" json:"user_id"`
	ReceiverId int       `gorm:"column:receiver_id;default:0;NOT NULL" json:"receiver_id"`
	IsTop      int       `gorm:"column:is_top;default:0;NOT NULL" json:"is_top"`
	IsDisturb  int       `gorm:"column:is_disturb;default:0;NOT NULL" json:"is_disturb"`
	IsDelete   int       `gorm:"column:is_delete;default:0;NOT NULL" json:"is_delete"`
	IsBot      int       `gorm:"column:is_bot;default:0;NOT NULL" json:"is_bot"`
	CreatedAt  time.Time `gorm:"column:created_at;NOT NULL" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;NOT NULL" json:"updated_at"`
}

func (Chat) TableName() string {
	return "chats"
}
