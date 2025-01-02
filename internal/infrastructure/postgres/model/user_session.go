// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package model

import "time"

type UserSession struct {
	Id          int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	UserId      int       `gorm:"column:user_id;NOT NULL" json:"user_id"`
	AccessToken string    `gorm:"column:access_token;NOT NULL" json:"access_token"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
	Logout      bool      `gorm:"column:is_logout;default false" json:"is_logout"`
	LogoutAt    time.Time `gorm:"column:logout_at;"`
	UserIp      string    `gorm:"column:user_ip;" json:"user_ip"`
	UserAgent   string    `gorm:"column:user_agent;" json:"user_agent"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func (UserSession) TableName() string {
	return "user_sessions"
}
