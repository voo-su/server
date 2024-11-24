package model

import "time"

type PushToken struct {
	Id          int64     `gorm:"column:id;primaryKey;autoIncrement"`
	UserId      int64     `gorm:"column:user_id;index;NOT NULL"`
	Platform    string    `gorm:"column:platform;NOT NULL"`
	Token       string    `gorm:"column:token"`
	WebEndpoint string    `gorm:"column:web_endpoint"`
	WebP256dh   string    `gorm:"column:web_p256dh"`
	WebAuth     string    `gorm:"column:web_auth"`
	IsActive    bool      `gorm:"column:is_active;default:true"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (PushToken) TableName() string {
	return "push_tokens"
}
