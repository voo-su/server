package model

import "time"

type MessageVote struct {
	Id           int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	RecordId     int       `gorm:"column:record_id;default:0;NOT NULL" json:"record_id"`
	UserId       int       `gorm:"column:user_id;default:0;NOT NULL" json:"user_id"`
	Title        string    `gorm:"column:title;NOT NULL" json:"title"`
	AnswerMode   int       `gorm:"column:answer_mode;default:0;NOT NULL" json:"answer_mode"`
	AnswerOption string    `gorm:"column:answer_option;NOT NULL" json:"answer_option"`
	AnswerNum    int       `gorm:"column:answer_num;default:0;NOT NULL" json:"answer_num"`
	AnsweredNum  int       `gorm:"column:answered_num;default:0;NOT NULL" json:"answered_num"`
	IsAnonymous  int       `gorm:"column:is_anonymous;default:0;NOT NULL" json:"is_anonymous"`
	Status       int       `gorm:"column:status;default:0;NOT NULL" json:"status"`
	CreatedAt    time.Time `gorm:"column:created_at;NOT NULL" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;NOT NULL" json:"updated_at"`
}

func (MessageVote) TableName() string {
	return "message_votes"
}
