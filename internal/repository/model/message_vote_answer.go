package model

import "time"

type MessageVoteAnswers struct {
	Id        int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	VoteId    int       `gorm:"column:vote_id;default:0;NOT NULL" json:"vote_id"`
	UserId    int       `gorm:"column:user_id;default:0;NOT NULL" json:"user_id"`
	Option    string    `gorm:"column:option;NOT NULL" json:"option"`
	CreatedAt time.Time `gorm:"column:created_at;NOT NULL" json:"created_at"`
}
