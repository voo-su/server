package model

import "time"

type ProjectMember struct {
	Id        int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	ProjectId int       `gorm:"column:project_id" json:"project_id"`
	UserId    int       `gorm:"column:user_id" json:"user_id"`
	CreatedBy int       `gorm:"column:created_by" json:"created_by"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

func (ProjectMember) TableName() string {
	return "project_members"
}

type ProjectMemberItem struct {
	Id       int64  `json:"id"`
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
}
