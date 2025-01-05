package entity

import "voo.su/internal/infrastructure/postgres/model"

type ProjectItem struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ProjectMemberItem struct {
	Id       int64  `json:"id"`
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
}

type ProjectTaskDetailWithMember struct {
	model.ProjectTask
	AssignerId       int64  `gorm:"column:assigner_id"`
	AssignerUsername string `gorm:"column:assigner_username"`
	AssignerName     string `gorm:"column:assigner_name"`
	AssignerSurname  string `gorm:"column:assigner_surname"`
	ExecutorId       int64  `gorm:"column:executor_id"`
	ExecutorUsername string `gorm:"column:executor_username"`
	ExecutorName     string `gorm:"column:executor_name"`
	ExecutorSurname  string `gorm:"column:executor_surname"`
}
