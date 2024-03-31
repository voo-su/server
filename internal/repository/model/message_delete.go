package model

import "time"

type MessageDelete struct {
	Id        int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	RecordId  int       `gorm:"column:record_id;default:0;NOT NULL" json:"record_id"`
	UserId    int       `gorm:"column:user_id;default:0;NOT NULL" json:"user_id"`
	CreatedAt time.Time `gorm:"column:created_at;NOT NULL" json:"created_at"`
}
