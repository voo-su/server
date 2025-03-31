package model

import "time"

type MessageLogin struct {
	Id        int       `gorm:"primaryKey"`
	MessageId int       `gorm:"column:message_id;NOT NULL"`
	IpAddress string    `gorm:"column:ip_address"`
	UserAgent string    `gorm:"column:user_agent"`
	Address   *string   `gorm:"column:address;NOT NULL"`
	UserId    int       `gorm:"column:user_id;NOT NULL"`
	CreatedAt time.Time `gorm:"column:created_at;NOT NULL"`
}

func (m MessageLogin) TableName() string {
	return "message_login"
}
