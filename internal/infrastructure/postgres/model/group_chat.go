// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package model

import (
	"time"
)

type GroupChat struct {
	Id          int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Type        int       `gorm:"column:type;default:1;NOT NULL" json:"type"`
	CreatorId   int       `gorm:"column:creator_id;default:0;NOT NULL" json:"creator_id"`
	Name        string    `gorm:"column:group_name;NOT NULL" json:"group_name"`
	Description string    `gorm:"column:description;NOT NULL" json:"description"`
	IsDismiss   int       `gorm:"column:is_dismiss;default:0;NOT NULL" json:"is_dismiss"`
	Avatar      string    `gorm:"column:avatar;NOT NULL" json:"avatar"`
	MaxNum      int       `gorm:"column:max_num;default:200;NOT NULL" json:"max_num"`
	IsOvert     int       `gorm:"column:is_overt;default:0;NOT NULL" json:"is_overt"`
	IsMute      int       `gorm:"column:is_mute;default:0;NOT NULL" json:"is_mute"`
	CreatedAt   time.Time `gorm:"column:created_at;NOT NULL" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;NOT NULL" json:"updated_at"`
}

func (GroupChat) TableName() string {
	return "group_chats"
}
