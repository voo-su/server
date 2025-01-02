// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package entity

import (
	"time"
	"voo.su/internal/constant"
	"voo.su/pkg/locale"
)

type SearchChat struct {
	Id          int       `json:"id" `
	DialogType  int       `json:"dialog_type" `
	ReceiverId  int       `json:"receiver_id" `
	IsDelete    int       `json:"is_delete"`
	IsTop       int       `json:"is_top"`
	IsBot       int       `json:"is_bot"`
	IsDisturb   int       `json:"is_disturb"`
	UserAvatar  string    `json:"user_avatar"`
	Username    string    `json:"username"`
	Name        string    `json:"name"`
	Surname     string    `json:"surname"`
	GroupName   string    `json:"group_name"`
	GroupAvatar string    `json:"group_avatar"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func GetChatMsgTypeMapping(locale locale.ILocale, msgType int) string {
	var ChatMsgTypeMapping = map[int]string{
		constant.ChatMsgTypeImage:                 locale.Localize("photo"),
		constant.ChatMsgTypeAudio:                 locale.Localize("audio_recording"),
		constant.ChatMsgTypeVideo:                 locale.Localize("video"),
		constant.ChatMsgTypeFile:                  locale.Localize("file"),
		constant.ChatMsgTypeLocation:              locale.Localize("location"),
		constant.ChatMsgTypeCard:                  locale.Localize("contact_info"),
		constant.ChatMsgTypeForward:               locale.Localize("forwarded_message"),
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
