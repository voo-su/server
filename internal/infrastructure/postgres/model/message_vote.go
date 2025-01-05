package model

import "time"

type MessageVote struct {
	Id           int       `gorm:"primaryKey"`
	RecordId     int       `gorm:"column:record_id;default:0;NOT NULL"`
	UserId       int       `gorm:"column:user_id;default:0;NOT NULL"`
	Title        string    `gorm:"column:title;NOT NULL"`
	AnswerMode   int       `gorm:"column:answer_mode;default:0;NOT NULL"`
	AnswerOption string    `gorm:"column:answer_option;NOT NULL"`
	AnswerNum    int       `gorm:"column:answer_num;default:0;NOT NULL"`
	AnsweredNum  int       `gorm:"column:answered_num;default:0;NOT NULL"`
	IsAnonymous  int       `gorm:"column:is_anonymous;default:0;NOT NULL"`
	Status       int       `gorm:"column:status;default:0;NOT NULL"`
	CreatedAt    time.Time `gorm:"column:created_at;NOT NULL"`
	UpdatedAt    time.Time `gorm:"column:updated_at;NOT NULL"`
}

func (MessageVote) TableName() string {
	return "message_votes"
}
