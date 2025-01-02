// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package model

import "time"

type StickerUser struct {
	Id         int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	UserId     int       `gorm:"column:user_id;NOT NULL" json:"user_id"`
	StickerIds string    `gorm:"column:sticker_ids;NOT NULL" json:"sticker_ids"`
	CreatedAt  time.Time `gorm:"column:created_at;NOT NULL" json:"created_at"`
}

func (StickerUser) TableName() string {
	return "sticker_users"
}
