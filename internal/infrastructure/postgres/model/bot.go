// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package model

import "time"

type Bot struct {
	Id          int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	UserId      int       `gorm:"column:user_id;default:0;NOT NULL" json:"user_id"`
	Token       string    `gorm:"column:token;unique;NOT NULL" json:"token"`
	Name        string    `gorm:"column:name;NOT NULL" json:"name"`
	Description string    `gorm:"column:description;NOT NULL" json:"description"`
	Avatar      string    `gorm:"column:avatar;NOT NULL" json:"avatar"`
	CreatedAt   time.Time `gorm:"column:created_at;NOT NULL" json:"created_at"`
	CreatorId   int       `gorm:"column:creator_id;default:NULL" json:"creator_id"`
}

func (Bot) TableName() string {
	return "bots"
}
