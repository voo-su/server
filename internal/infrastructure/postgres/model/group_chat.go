package model

import (
	"time"
)

type GroupChat struct {
	Id          int       `gorm:"primaryKey"`
	Type        int       `gorm:"column:type;DEFAULT:1;NOT NULL"`
	CreatorId   int       `gorm:"column:creator_id;DEFAULT:0;NOT NULL"`
	Name        string    `gorm:"column:group_name;NOT NULL"`
	Description string    `gorm:"column:description;NOT NULL"`
	IsDismiss   int       `gorm:"column:is_dismiss;DEFAULT:0;NOT NULL"`
	Avatar      string    `gorm:"column:avatar;NOT NULL"`
	MaxNum      int       `gorm:"column:max_num;DEFAULT:200;NOT NULL"`
	IsOvert     int       `gorm:"column:is_overt;DEFAULT:0;NOT NULL"`
	IsMute      int       `gorm:"column:is_mute;DEFAULT:0;NOT NULL"`
	CreatedAt   time.Time `gorm:"column:created_at;NOT NULL"`
	UpdatedAt   time.Time `gorm:"column:updated_at;NOT NULL"`
}

func (GroupChat) TableName() string {
	return "group_chats"
}
