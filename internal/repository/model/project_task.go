package model

import "time"

type ProjectTask struct {
	Id          int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	ProjectId   int64     `gorm:"column:project_id" json:"project_id"`
	TaskType    int       `gorm:"column:task_type" json:"task_type"`
	Title       string    `gorm:"column:title" json:"title"`
	Description string    `gorm:"column:description" json:"description"`
	CreatedBy   int       `gorm:"column:created_by" json:"created_by"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
}

func (ProjectTask) TableName() string {
	return "project_tasks"
}
