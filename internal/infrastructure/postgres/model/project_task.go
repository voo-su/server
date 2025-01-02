// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package model

import "time"

type ProjectTask struct {
	Id          int64     `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	ProjectId   int64     `gorm:"column:project_id" json:"project_id"`
	TypeId      int       `gorm:"column:type_id" json:"type_id"`
	Title       string    `gorm:"column:title" json:"title"`
	Description string    `gorm:"column:description" json:"description"`
	AssignerId  int       `gorm:"column:assigner_id" json:"assigner_id"`
	ExecutorId  int       `gorm:"column:executor_id" json:"executor_id"`
	CreatedBy   int       `gorm:"column:created_by" json:"created_by"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
}

func (ProjectTask) TableName() string {
	return "project_tasks"
}

type ProjectTaskDetailWithMember struct {
	ProjectTask
	AssignerId       int64  `gorm:"column:assigner_id" json:"assigner_id"`
	AssignerUsername string `gorm:"column:assigner_username" json:"assigner_username"`
	ExecutorId       int64  `gorm:"column:executor_id" json:"executor_id"`
	ExecutorUsername string `gorm:"column:executor_username" json:"executor_username"`
}
