package model

import "time"

type Chat struct {
	Id                 int       `gorm:"primaryKey"`
	ChatType           int       `gorm:"column:chat_type;DEFAULT:1;NOT NULL"`
	UserId             int       `gorm:"column:user_id;DEFAULT:0;NOT NULL"`
	ReceiverId         int       `gorm:"column:receiver_id;DEFAULT:0;NOT NULL"`
	IsTop              int       `gorm:"column:is_top;DEFAULT:0;NOT NULL"`
	NotifyMuteUntil    int32     `gorm:"column:notify_mute_until;DEFAULT:0;NOT NULL"`
	NotifyShowPreviews bool      `gorm:"column:notify_show_previews;DEFAULT:true;NOT NULL"`
	NotifySilent       bool      `gorm:"column:notify_silent;DEFAULT:FALSE;NOT NULL"`
	IsDisturb          int       `gorm:"column:is_disturb;DEFAULT:0;NOT NULL"`
	IsDelete           int       `gorm:"column:is_delete;DEFAULT:0;NOT NULL"`
	IsBot              int       `gorm:"column:is_bot;DEFAULT:0;NOT NULL"`
	CreatedAt          time.Time `gorm:"column:created_at;NOT NULL"`
	UpdatedAt          time.Time `gorm:"column:updated_at;NOT NULL"`
}

func (Chat) TableName() string {
	return "chats"
}
