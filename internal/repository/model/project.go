package model

import "time"

type Project struct {
	Id        int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	CreatedBy int       `gorm:"column:created_by" json:"created_by"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

func (Project) TableName() string {
	return "projects"
}
