package model

import "time"

type Bot struct {
	Id          int       `gorm:"primaryKey"`
	UserId      int       `gorm:"column:user_id;default:0;NOT NULL"`
	Token       string    `gorm:"column:token;unique;NOT NULL"`
	Name        string    `gorm:"column:name;NOT NULL"`
	Description string    `gorm:"column:description;NOT NULL"`
	Avatar      string    `gorm:"column:avatar;NOT NULL"`
	CreatedAt   time.Time `gorm:"column:created_at;NOT NULL"`
	CreatorId   int       `gorm:"column:creator_id;default:NULL"`
}

func (Bot) TableName() string {
	return "bots"
}
