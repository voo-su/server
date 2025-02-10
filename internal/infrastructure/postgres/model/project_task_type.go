package model

import (
	"github.com/google/uuid"
	"time"
)

type ProjectTaskType struct {
	Id        int64     `gorm:"primaryKey"`
	ProjectId uuid.UUID `gorm:"column:project_id"`
	Title     string    `gorm:"column:title"`
	CreatedBy int       `gorm:"column:created_by"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (ProjectTaskType) TableName() string {
	return "project_task_types"
}
