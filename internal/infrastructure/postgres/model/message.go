// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package model

import "time"

type Message struct {
	Id         int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	MsgId      string    `gorm:"column:msg_id;NOT NULL" json:"msg_id"`
	Sequence   int64     `gorm:"column:sequence;default:0;NOT NULL" json:"sequence"`
	DialogType int       `gorm:"column:dialog_type;default:1;NOT NULL" json:"dialog_type"`
	MsgType    int       `gorm:"column:msg_type;default:0;NOT NULL" json:"msg_type"`
	UserId     int       `gorm:"column:user_id;default:0;NOT NULL" json:"user_id"`
	ReceiverId int       `gorm:"column:receiver_id;default:0;NOT NULL" json:"receiver_id"`
	IsRevoke   int       `gorm:"column:is_revoke;default:0;NOT NULL" json:"is_revoke"`
	IsMark     int       `gorm:"column:is_mark;default:0;NOT NULL" json:"is_mark"`
	IsRead     int       `gorm:"column:is_read;default:0;NOT NULL" json:"is_read"`
	QuoteId    string    `gorm:"column:quote_id;NOT NULL" json:"quote_id"`
	Content    string    `gorm:"column:content" json:"content"`
	Extra      string    `gorm:"column:extra;default:{}" json:"extra"`
	CreatedAt  time.Time `gorm:"column:created_at;NOT NULL" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;NOT NULL" json:"updated_at"`
}

func (t Message) TableName() string {
	return "messages"
}
