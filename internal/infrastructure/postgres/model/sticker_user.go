package model

import "time"

type StickerUser struct {
	Id         int       `gorm:"primaryKey"`
	UserId     int       `gorm:"column:user_id;NOT NULL"`
	StickerIds string    `gorm:"column:sticker_ids;NOT NULL"`
	CreatedAt  time.Time `gorm:"column:created_at;NOT NULL"`
}

func (StickerUser) TableName() string {
	return "sticker_users"
}
