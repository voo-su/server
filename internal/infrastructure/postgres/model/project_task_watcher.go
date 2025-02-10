package model

import (
	"github.com/google/uuid"
	"time"
)

type ProjectTaskWatcher struct {
	Id        int64     `gorm:"primaryKey"`
	TaskId    uuid.UUID `gorm:"column:task_id"`
	MemberId  int       `gorm:"column:member_id"`
	CreatedBy int       `gorm:"column:created_by"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (ProjectTaskWatcher) TableName() string {
	return "project_task_watchers"
}
