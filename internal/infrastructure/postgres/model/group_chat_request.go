// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package model

import "time"

type GroupChatRequest struct {
	Id        int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	GroupId   int       `gorm:"column:group_id;default:0;NOT NULL" json:"group_id"`
	UserId    int       `gorm:"column:user_id;default:0;NOT NULL" json:"user_id"`
	Status    int       `gorm:"column:status;default:1;NOT NULL" json:"status"`
	CreatedAt time.Time `gorm:"column:created_at;NOT NULL" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;NOT NULL" json:"updated_at"`
}

func (GroupChatRequest) TableName() string {
	return "group_chat_requests"
}
