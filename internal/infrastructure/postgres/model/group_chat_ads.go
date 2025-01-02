// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package model

import (
	"database/sql"
	"time"
)

type GroupChatAds struct {
	Id           int          `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	GroupId      int          `gorm:"column:group_id;default:0;NOT NULL" json:"group_id"`
	CreatorId    int          `gorm:"column:creator_id;default:0;NOT NULL" json:"creator_id"`
	Title        string       `gorm:"column:title;NOT NULL" json:"title"`
	Content      string       `gorm:"column:content;NOT NULL" json:"content"`
	ConfirmUsers string       `gorm:"column:confirm_users" json:"confirm_users"`
	IsDelete     int          `gorm:"column:is_delete;default:0;NOT NULL" json:"is_delete"`
	IsTop        int          `gorm:"column:is_top;default:0;NOT NULL" json:"is_top"`
	IsConfirm    int          `gorm:"column:is_confirm;default:0;NOT NULL" json:"is_confirm"`
	CreatedAt    time.Time    `gorm:"column:created_at;NOT NULL" json:"created_at"`
	UpdatedAt    time.Time    `gorm:"column:updated_at;NOT NULL" json:"updated_at"`
	DeletedAt    sql.NullTime `gorm:"column:deleted_at" json:"deleted_at"`
}

func (GroupChatAds) TableName() string {
	return "group_chat_ads"
}
