package model

import "time"

type ProjectTaskComment struct {
	Id        int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	TaskId    int       `gorm:"column:task_id" json:"task_id"`
	Text      string    `gorm:"column:comment_text" json:"comment_text"`
	CreatedBy int       `gorm:"column:created_by" json:"created_by"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

func (ProjectTaskComment) TableName() string {
	return "project_task_comments"
}
