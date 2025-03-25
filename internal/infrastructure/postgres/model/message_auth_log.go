package model

import "time"

type MessageAuthLog struct {
	Id        int       `gorm:"primaryKey"`
	IpAddress string    `gorm:"column:ip_address"`
	UserAgent string    `gorm:"column:user_agent"`
	UserId    int       `gorm:"column:user_id;NOT NULL"`
	CreatedAt time.Time `gorm:"column:created_at;NOT NULL"`
}

func (m MessageAuthLog) TableName() string {
	return "message_auth_logs"
}
