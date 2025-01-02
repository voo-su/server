// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package model

import "time"

type ProjectTaskWatcher struct {
	Id        int64     `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	TaskId    int       `gorm:"column:task_id" json:"task_id"`
	MemberId  int       `gorm:"column:member_id" json:"member_id"`
	CreatedBy int       `gorm:"column:created_by" json:"created_by"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

func (ProjectTaskWatcher) TableName() string {
	return "project_task_watchers"
}
