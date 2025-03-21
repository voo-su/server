package model

import "time"

type Logger struct {
	LogMessage string     `ch:"log_message"`
	CreatedAt  *time.Time `ch:"created_at"`
}

func (Logger) TableName() string {
	return "loggers"
}
