package model

import (
	"time"
)

type MessageLocation struct {
	Id          int       `gorm:"primaryKey"`
	MessageId   int       `gorm:"column:message_id;NOT NULL"`
	Longitude   string    `json:"longitude"`
	Latitude    string    `json:"latitude"`
	Description string    `json:"description"`
	CreatedAt   time.Time `gorm:"column:created_at;NOT NULL"`
}

func (m MessageLocation) TableName() string {
	return "message_location"
}
