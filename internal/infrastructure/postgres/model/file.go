package model

import (
	"github.com/google/uuid"
	"time"
)

type File struct {
	Id           uuid.UUID `gorm:"column:id;type:uuid;DEFAULT:gen_random_uuid()"`
	OriginalName string    `gorm:"column:original_name;NOT NULL"`
	ObjectName   string    `gorm:"column:object_name;NOT NULL"`
	Size         int       `gorm:"column:size;NOT NULL"`
	MimeType     string    `gorm:"column:mime_type;NOT NULL"`
	CreatedBy    int       `gorm:"column:created_by;NOT NULL"`
	CreatedAt    time.Time `gorm:"column:created_at;NOT NULL"`
}

func (File) TableName() string {
	return "files"
}
