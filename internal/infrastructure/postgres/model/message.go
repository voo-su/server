package model

import (
	"github.com/google/uuid"
	"time"
)

type Message struct {
	Id         int        `gorm:"primaryKey"`
	MsgId      string     `gorm:"column:msg_id;NOT NULL"`
	Sequence   int64      `gorm:"column:sequence;DEFAULT:0;NOT NULL"`
	ChatType   int        `gorm:"column:chat_type;DEFAULT:1;NOT NULL"`
	MsgType    int        `gorm:"column:msg_type;DEFAULT:0;NOT NULL"`
	UserId     int        `gorm:"column:user_id;DEFAULT:0;NOT NULL"`
	ReceiverId int        `gorm:"column:receiver_id;DEFAULT:0;NOT NULL"`
	IsRevoke   int        `gorm:"column:is_revoke;DEFAULT:0;NOT NULL"`
	IsMark     int        `gorm:"column:is_mark;DEFAULT:0;NOT NULL"`
	IsRead     int        `gorm:"column:is_read;DEFAULT:0;NOT NULL"`
	QuoteId    string     `gorm:"column:quote_id;NOT NULL"`
	Content    string     `gorm:"column:content"`
	Extra      string     `gorm:"column:extra;DEFAULT:{}"`
	FileId     *uuid.UUID `gorm:"column:file_id;DEFAULT:NULL"`
	ReplyTo    int        `gorm:"column:reply_to;DEFAULT:NULL"`
	CreatedAt  time.Time  `gorm:"column:created_at;NOT NULL"`
	UpdatedAt  time.Time  `gorm:"column:updated_at;NOT NULL"`
}

func (t Message) TableName() string {
	return "messages"
}
