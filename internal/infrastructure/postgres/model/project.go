package model

import (
	"github.com/google/uuid"
	"time"
)

type Project struct {
	Id        uuid.UUID `gorm:"column:id;type:uuid;DEFAULT:gen_random_uuid()"`
	Name      string    `gorm:"column:name"`
	CreatedBy int       `gorm:"column:created_by"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (Project) TableName() string {
	return "projects"
}
