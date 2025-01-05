package model

import "time"

type Project struct {
	Id        int       `gorm:"primaryKey"`
	Name      string    `gorm:"column:name"`
	CreatedBy int       `gorm:"column:created_by"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (Project) TableName() string {
	return "projects"
}
