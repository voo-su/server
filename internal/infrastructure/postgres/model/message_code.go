package model

import (
	"time"
)

type MessageCode struct {
	Id        int       `gorm:"primaryKey"`
	MessageId int       `gorm:"column:message_id;NOT NULL"`
	Lang      string    `gorm:"column:lang"`
	Code      string    `gorm:"column:code"`
	CreatedAt time.Time `gorm:"column:created_at;NOT NULL"`
}

func (m MessageCode) TableName() string {
	return "message_code"
}
