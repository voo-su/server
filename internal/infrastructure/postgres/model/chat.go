package model

import "time"

type Chat struct {
	Id         int       `gorm:"primaryKey"`
	ChatType   int       `gorm:"column:chat_type;default:1;NOT NULL"`
	UserId     int       `gorm:"column:user_id;default:0;NOT NULL"`
	ReceiverId int       `gorm:"column:receiver_id;default:0;NOT NULL"`
	IsTop      int       `gorm:"column:is_top;default:0;NOT NULL"`
	IsDisturb  int       `gorm:"column:is_disturb;default:0;NOT NULL"`
	IsDelete   int       `gorm:"column:is_delete;default:0;NOT NULL"`
	IsBot      int       `gorm:"column:is_bot;default:0;NOT NULL"`
	CreatedAt  time.Time `gorm:"column:created_at;NOT NULL"`
	UpdatedAt  time.Time `gorm:"column:updated_at;NOT NULL"`
}

func (Chat) TableName() string {
	return "chats"
}
