package model

import (
	"database/sql"
	"time"
)

type GroupChatAds struct {
	Id           int          `gorm:"primaryKey"`
	GroupId      int          `gorm:"column:group_id;default:0;NOT NULL"`
	CreatorId    int          `gorm:"column:creator_id;default:0;NOT NULL"`
	Title        string       `gorm:"column:title;NOT NULL"`
	Content      string       `gorm:"column:content;NOT NULL"`
	ConfirmUsers string       `gorm:"column:confirm_users"`
	IsDelete     int          `gorm:"column:is_delete;default:0;NOT NULL"`
	IsTop        int          `gorm:"column:is_top;default:0;NOT NULL"`
	IsConfirm    int          `gorm:"column:is_confirm;default:0;NOT NULL"`
	CreatedAt    time.Time    `gorm:"column:created_at;NOT NULL"`
	UpdatedAt    time.Time    `gorm:"column:updated_at;NOT NULL"`
	DeletedAt    sql.NullTime `gorm:"column:deleted_at"`
}

func (GroupChatAds) TableName() string {
	return "group_chat_ads"
}
