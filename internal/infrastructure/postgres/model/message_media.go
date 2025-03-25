package model

import (
	"github.com/google/uuid"
	"time"
)

type MessageMedia struct {
	Id        int       `gorm:"primaryKey"`
	MessageId string    `gorm:"column:message_id;NOT NULL"`
	FileId    uuid.UUID `gorm:"column:file_id"`
	Duration  int       `gorm:"column:duration"`
	CreatedAt time.Time `gorm:"column:created_at;NOT NULL"`
}

func (m MessageMedia) TableName() string {
	return "message_media"
}
