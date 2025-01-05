package model

import "time"

type ProjectTask struct {
	Id          int64     `gorm:"primaryKey"`
	ProjectId   int64     `gorm:"column:project_id"`
	TypeId      int       `gorm:"column:type_id"`
	Title       string    `gorm:"column:title"`
	Description string    `gorm:"column:description"`
	AssignerId  int       `gorm:"column:assigner_id"`
	ExecutorId  int       `gorm:"column:executor_id"`
	CreatedBy   int       `gorm:"column:created_by"`
	CreatedAt   time.Time `gorm:"column:created_at"`
}

func (ProjectTask) TableName() string {
	return "project_tasks"
}
