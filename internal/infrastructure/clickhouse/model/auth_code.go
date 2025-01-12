package model

import "time"

type AuthCode struct {
	Email        string    `ch:"email"`
	Code         string    `ch:"code"`
	Token        string    `ch:"token"`
	ErrorMessage string    `ch:"error_message"`
	CreatedAt    time.Time `ch:"created_at"`
}

func (AuthCode) TableName() string {
	return "auth_codes"
}
