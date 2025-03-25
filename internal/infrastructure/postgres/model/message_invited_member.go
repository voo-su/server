package model

import (
	"time"
)

type MessageInvitedMember struct {
	Id        int       `gorm:"primaryKey"`
	MessageId string    `gorm:"column:message_id;NOT NULL"`
	UserId    int       `gorm:"column:user_id;NOT NULL"`
	CreatedAt time.Time `gorm:"column:created_at;NOT NULL"`
}

func (m MessageInvitedMember) TableName() string {
	return "message_invited_members"
}
