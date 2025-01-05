package model

import "time"

type Message struct {
	Id         int       `gorm:"primaryKey"`
	MsgId      string    `gorm:"column:msg_id;NOT NULL"`
	Sequence   int64     `gorm:"column:sequence;default:0;NOT NULL"`
	DialogType int       `gorm:"column:dialog_type;default:1;NOT NULL"`
	MsgType    int       `gorm:"column:msg_type;default:0;NOT NULL"`
	UserId     int       `gorm:"column:user_id;default:0;NOT NULL"`
	ReceiverId int       `gorm:"column:receiver_id;default:0;NOT NULL"`
	IsRevoke   int       `gorm:"column:is_revoke;default:0;NOT NULL"`
	IsMark     int       `gorm:"column:is_mark;default:0;NOT NULL"`
	IsRead     int       `gorm:"column:is_read;default:0;NOT NULL"`
	QuoteId    string    `gorm:"column:quote_id;NOT NULL"`
	Content    string    `gorm:"column:content"`
	Extra      string    `gorm:"column:extra;default:{}"`
	CreatedAt  time.Time `gorm:"column:created_at;NOT NULL"`
	UpdatedAt  time.Time `gorm:"column:updated_at;NOT NULL"`
}

func (t Message) TableName() string {
	return "messages"
}
