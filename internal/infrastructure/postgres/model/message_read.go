package model

import "time"

type MessageRead struct {
	Id         int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	MsgId      string    `gorm:"column:msg_id;NOT NULL" json:"msg_id"`
	UserId     int       `gorm:"column:user_id;default:0;NOT NULL" json:"user_id"`
	ReceiverId int       `gorm:"column:receiver_id;default:0;NOT NULL" json:"receiver_id"`
	CreatedAt  time.Time `gorm:"column:created_at;NOT NULL" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;NOT NULL" json:"updated_at"`
}

func (m MessageRead) TableName() string {
	return "message_read"
}
