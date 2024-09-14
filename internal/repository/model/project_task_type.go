package model

import "time"

type ProjectTaskType struct {
	Id        int64     `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Title     string    `gorm:"column:title" json:"title"`
	CreatedBy int       `gorm:"column:created_by" json:"created_by"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

func (ProjectTaskType) TableName() string {
	return "project_task_types"
}
