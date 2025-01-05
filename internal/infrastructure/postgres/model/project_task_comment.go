package model

import "time"

type ProjectTaskComment struct {
	Id        int64     `gorm:"primaryKey"`
	TaskId    int64     `gorm:"column:task_id"`
	Comment   string    `gorm:"column:comment_text"`
	CreatedBy int       `gorm:"column:created_by"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (ProjectTaskComment) TableName() string {
	return "project_task_comments"
}
