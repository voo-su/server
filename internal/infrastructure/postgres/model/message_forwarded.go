package model

import "time"

type MessageForwarded struct {
	Id                int       `gorm:"primaryKey"`
	OriginalMessageId int       `gorm:"column:original_message_id;NOT NULL"`
	NewMessageId      int       `gorm:"column:new_message_id;NOT NULL"`
	UserId            int       `gorm:"column:user_id;NOT NULL"`
	CreatedAt         time.Time `gorm:"column:created_at;NOT NULL"`
}

func (m MessageForwarded) TableName() string {
	return "message_forwarded"
}
