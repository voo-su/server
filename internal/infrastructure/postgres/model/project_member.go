package model

import (
	"github.com/google/uuid"
	"time"
)

type ProjectMember struct {
	Id        int       `gorm:"primaryKey"`
	ProjectId uuid.UUID `gorm:"column:project_id"`
	UserId    int       `gorm:"column:user_id"`
	CreatedBy int       `gorm:"column:created_by"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (ProjectMember) TableName() string {
	return "project_members"
}
