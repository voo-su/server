package model

import (
	"time"
)

type Message struct {
	Id         int       `gorm:"primaryKey"`
	Sequence   int64     `gorm:"column:sequence;DEFAULT:0;NOT NULL"`
	ChatType   int       `gorm:"column:chat_type;DEFAULT:1;NOT NULL"`
	MsgType    int       `gorm:"column:msg_type;DEFAULT:0;NOT NULL"`
	UserId     int       `gorm:"column:user_id;DEFAULT:0;NOT NULL"`
	ReceiverId int       `gorm:"column:receiver_id;DEFAULT:0;NOT NULL"`
	IsRevoke   int       `gorm:"column:is_revoke;DEFAULT:0;NOT NULL"`
	IsMark     int       `gorm:"column:is_mark;DEFAULT:0;NOT NULL"`
	IsRead     int       `gorm:"column:is_read;DEFAULT:0;NOT NULL"`
	Content    string    `gorm:"column:content"`
	Extra      string    `gorm:"column:extra;DEFAULT:{}"`
	ReplyTo    int       `gorm:"column:reply_to;DEFAULT:NULL"`
	CreatedAt  time.Time `gorm:"column:created_at;NOT NULL"`
	UpdatedAt  time.Time `gorm:"column:updated_at;NOT NULL"`
	//QuoteId    string    `gorm:"column:quote_id;NOT NULL"`
}

func (t Message) TableName() string {
	return "messages"
}
