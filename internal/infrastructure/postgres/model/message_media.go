package model

import (
	"github.com/google/uuid"
	"time"
)

type MessageMedia struct {
	Id        int       `gorm:"primaryKey"`
	MessageId int       `gorm:"column:message_id;NOT NULL"`
	FileId    uuid.UUID `gorm:"column:file_id"`
	Drive     int       `gorm:"column:drive"`
	Duration  int       `gorm:"column:duration"`
	Url       string    `gorm:"column:url"`
	Name      string    `gorm:"column:name"`
	Size      int       `gorm:"column:size"`
	Cover     string    `gorm:"column:cover"`
	MimeType  string    `gorm:"column:mime_type"`
	Width     int32     `gorm:"column:width"`
	Height    int32     `gorm:"column:height"`
	CreatedAt time.Time `gorm:"column:created_at;NOT NULL"`
}

func (m MessageMedia) TableName() string {
	return "message_media"
}
