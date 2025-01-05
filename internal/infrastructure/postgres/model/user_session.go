package model

import "time"

type UserSession struct {
	Id          int       `gorm:"primaryKey"`
	UserId      int       `gorm:"column:user_id;NOT NULL"`
	AccessToken string    `gorm:"column:access_token;NOT NULL"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
	Logout      bool      `gorm:"column:is_logout;default false"`
	LogoutAt    time.Time `gorm:"column:logout_at;"`
	UserIp      string    `gorm:"column:user_ip;"`
	UserAgent   string    `gorm:"column:user_agent;"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func (UserSession) TableName() string {
	return "user_sessions"
}
