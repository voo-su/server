package model

import "time"

type MessageDelete struct {
	Id        int       `gorm:"primaryKey"`
	RecordId  int       `gorm:"column:record_id;default:0;NOT NULL"`
	UserId    int       `gorm:"column:user_id;default:0;NOT NULL"`
	CreatedAt time.Time `gorm:"column:created_at;NOT NULL"`
}

func (m MessageDelete) TableName() string {
	return "message_delete"
}
