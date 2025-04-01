package entity

import (
	"time"
	"voo.su/internal/constant"
	"voo.su/pkg/locale"
)

type SearchChat struct {
	Id                 int       `gorm:"id" `
	ChatType           int       `gorm:"chat_type" `
	ReceiverId         int       `gorm:"receiver_id" `
	IsDelete           int       `gorm:"is_delete"`
	IsTop              int       `gorm:"is_top"`
	IsBot              int       `gorm:"is_bot"`
	NotifyMuteUntil    int32     `gorm:"column:notify_mute_until"`
	NotifyShowPreviews bool      `gorm:"column:notify_show_previews"`
	NotifySilent       bool      `gorm:"column:notify_silent"`
	IsDisturb          int       `gorm:"is_disturb"`
	UserAvatar         string    `gorm:"user_avatar"`
	Username           string    `gorm:"username"`
	Name               string    `gorm:"name"`
	Surname            string    `gorm:"surname"`
	GroupName          string    `gorm:"group_name"`
	GroupAvatar        string    `gorm:"group_avatar"`
	UpdatedAt          time.Time `gorm:"updated_at"`
}

func GetChatMsgTypeMapping(locale locale.ILocale, msgType int) string {
	var ChatMsgTypeMapping = map[int]string{
		constant.ChatMsgTypeImage:                 locale.Localize("photo"),
		constant.ChatMsgTypeAudio:                 locale.Localize("audio_recording"),
		constant.ChatMsgTypeVideo:                 locale.Localize("video"),
		constant.ChatMsgTypeFile:                  locale.Localize("file"),
		constant.ChatMsgTypeLocation:              locale.Localize("location"),
		constant.ChatMsgTypeCard:                  locale.Localize("contact_info"),
		constant.ChatMsgTypeForwarded:             locale.Localize("forwarded_message"),
		constant.ChatMsgTypeLogin:                 locale.Localize("login_notification"),
		constant.ChatMsgTypeVote:                  locale.Localize("vote"),
		constant.ChatMsgTypeCode:                  locale.Localize("code"),
		constant.ChatMsgTypeMixed:                 locale.Localize("photos"),
		constant.ChatMsgSysText:                   locale.Localize("system_message"),
		constant.ChatMsgSysGroupCreate:            locale.Localize("group_creation"),
		constant.ChatMsgSysGroupMemberJoin:        locale.Localize("group_joining"),
		constant.ChatMsgSysGroupMemberQuit:        locale.Localize("group_exit"),
		constant.ChatMsgSysGroupMemberKicked:      locale.Localize("group_exclusion"),
		constant.ChatMsgSysGroupMessageRevoke:     locale.Localize("message_revoke"),
		constant.ChatMsgSysGroupDismissed:         locale.Localize("group_deletion"),
		constant.ChatMsgSysGroupMuted:             locale.Localize("group_notifications_off"),
		constant.ChatMsgSysGroupCancelMuted:       locale.Localize("group_notifications_on"),
		constant.ChatMsgSysGroupMemberMuted:       locale.Localize("participant_notifications_off"),
		constant.ChatMsgSysGroupMemberCancelMuted: locale.Localize("participant_notifications_on"),
		constant.ChatMsgSysGroupAds:               locale.Localize("group_announcement"),
	}

	if value, ok := ChatMsgTypeMapping[msgType]; ok {
		return value
	}

	return locale.Localize("unknown")
}

type Message struct {
	Id         int64  `json:"id"`
	ChatType   int    `json:"chat_type"`
	MsgType    int    `json:"msg_type"`
	ReceiverId int    `json:"receiver_id"`
	UserId     int    `json:"user_id"`
	Content    string `json:"content"`
	IsRead     bool   `json:"is_read"`
	CreatedAt  string `json:"created_at"`
}

type ConsumeMessage struct {
	UserIds []int   `json:"user_ids"`
	Message Message `json:"message"`
}

type SubscribeContent struct {
	Event string `json:"event"`
	Data  string `json:"data"`
}

type ConsumeChatKeyboard struct {
	SenderId   int `json:"sender_id"`
	ReceiverId int `json:"receiver_id"`
}
