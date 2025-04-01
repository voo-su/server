package model

import "time"

type MessageSystem struct {
	Id           int       `gorm:"primaryKey"`
	ChatType     int       `gorm:"column:chat_type;not null"`
	ReceiverId   int       `gorm:"column:receiver_id"`
	MessageId    int       `gorm:"column:message_id"`
	UserId       int       `gorm:"column:user_id"`
	TargetUserId int       `gorm:"column:target_user_id;DEFAULT:NULL"`
	EventType    int       `gorm:"column:event_type;not null"`
	OldValue     string    `gorm:"column:old_value;type:jsonb;DEFAULT:NULL"`
	NewValue     string    `gorm:"column:new_value;type:jsonb;DEFAULT:NULL"`
	Description  string    `gorm:"column:description;type:text;DEFAULT:NULL"`
	CreatedAt    time.Time `gorm:"column:created_at;not null"`
}

func (m MessageSystem) TableName() string {
	return "message_system"
}
