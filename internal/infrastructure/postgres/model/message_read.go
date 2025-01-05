package model

import "time"

type MessageRead struct {
	Id         int       `gorm:"primaryKey"`
	MsgId      string    `gorm:"column:msg_id;NOT NULL"`
	UserId     int       `gorm:"column:user_id;default:0;NOT NULL"`
	ReceiverId int       `gorm:"column:receiver_id;default:0;NOT NULL"`
	CreatedAt  time.Time `gorm:"column:created_at;NOT NULL"`
	UpdatedAt  time.Time `gorm:"column:updated_at;NOT NULL"`
}

func (m MessageRead) TableName() string {
	return "message_read"
}
