package model

import (
	"github.com/google/uuid"
	"time"
)

type ProjectTask struct {
	Id          uuid.UUID `gorm:"column:id;type:uuid;DEFAULT:gen_random_uuid()"`
	ProjectId   uuid.UUID `gorm:"column:project_id"`
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
