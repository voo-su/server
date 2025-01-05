package model

import "time"

type MessageVoteAnswer struct {
	Id        int       `gorm:"primaryKey"`
	VoteId    int       `gorm:"column:vote_id;default:0;NOT NULL"`
	UserId    int       `gorm:"column:user_id;default:0;NOT NULL"`
	Option    string    `gorm:"column:option;NOT NULL"`
	CreatedAt time.Time `gorm:"column:created_at;NOT NULL"`
}

func (MessageVoteAnswer) TableName() string {
	return "message_vote_answers"
}
