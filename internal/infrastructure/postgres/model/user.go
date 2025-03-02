package model

import "time"

type User struct {
	Id                      int       `gorm:"primaryKey"`
	Email                   string    `gorm:"column:email;NOT NULL"`
	Username                string    `gorm:"column:username;NOT NULL"`
	Name                    string    `gorm:"column:name;NOT NULL"`
	Surname                 string    `gorm:"column:surname;NOT NULL"`
	Avatar                  string    `gorm:"column:avatar;NOT NULL"`
	Gender                  int       `gorm:"column:gender;DEFAULT:0;NOT NULL"`
	About                   string    `gorm:"column:about;NOT NULL"`
	Birthday                string    `gorm:"column:birthday;NOT NULL"`
	IsBot                   int       `gorm:"column:is_bot;DEFAULT:0;NOT NULL"`
	NotifyChatsMuteUntil    int32     `gorm:"column:notify_chats_mute_until;DEFAULT:0;NOT NULL"`
	NotifyChatsShowPreviews bool      `gorm:"column:notify_chats_show_previews;DEFAULT:true;NOT NULL"`
	NotifyChatsSilent       bool      `gorm:"column:notify_chats_silent;DEFAULT:FALSE;NOT NULL"`
	NotifyGroupMuteUntil    int32     `gorm:"column:notify_group_mute_until;DEFAULT:0;NOT NULL"`
	NotifyGroupShowPreviews bool      `gorm:"column:notify_group_show_previews;DEFAULT:true;NOT NULL"`
	NotifyGroupSilent       bool      `gorm:"column:notify_group_silent;DEFAULT:FALSE;NOT NULL"`
	CreatedAt               time.Time `gorm:"column:created_at;NOT NULL"`
	UpdatedAt               time.Time `gorm:"column:updated_at;NOT NULL"`
}

func (User) TableName() string {
	return "users"
}
