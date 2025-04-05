package constant

const (
	ChatPrivateMode = 1
	ChatGroupMode   = 2
)

type RoomType string

const (
	RoomImGroup RoomType = "room_chat_group"
)

const (
	SubEventImMessage         = "sub.voo.message"
	SubEventImMessageKeyboard = "sub.voo.message.keyboard"
	SubEventImMessageRevoke   = "sub.voo.message.revoke"
	SubEventImMessageRead     = "sub.voo.message.read"
	SubEventContactStatus     = "sub.voo.contact.status"
	SubEventContactRequest    = "sub.voo.contact.request"
	SubEventGroupChatJoin     = "sub.voo.group.join"
	SubEventGroupChatRequest  = "sub.voo.group.request"

	PushEventImMessage         = "voo.message"
	PushEventImMessageKeyboard = "voo.message.keyboard"
	PushEventImMessageRead     = "voo.message.read"
	PushEventImMessageRevoke   = "voo.message.revoke"

	PushEventGroupChatRequest = "voo.group_chat.request"
)

const (
	ImChannelChat = "chat"

	ImTopicChat        = "im:message:chat:all"
	ImTopicChatPrivate = "im:message:chat:%s"
)

const (
	ChatTypePrivate = 1
	ChatTypeGroup   = 2
)

const (
	VoteAnswerModeSingleChoice = 0
)

// TODO iota

const (
	ChatMsgTypeUnknown = 0

	ChatMsgTypeText      = 1
	ChatMsgTypeCode      = 2
	ChatMsgTypeImage     = 3
	ChatMsgTypeAudio     = 4
	ChatMsgTypeVideo     = 5
	ChatMsgTypeFile      = 6
	ChatMsgTypeLocation  = 7
	ChatMsgTypeCard      = 8
	ChatMsgTypeForwarded = 9
	ChatMsgTypeLogin     = 10
	ChatMsgTypeVote      = 11

	ChatMsgSysText = 1000

	ChatMsgSysGroupCreate           = 1101
	ChatMsgSysGroupUserInvite       = 1102
	ChatMsgSysGroupUserLeave        = 1103
	ChatMsgSysGroupUserRemove       = 1104
	ChatMsgSysGroupMessageRevoke    = 1105
	ChatMsgSysGroupDismissed        = 1106
	ChatMsgSysGroupMuted            = 1107
	ChatMsgSysGroupCancelMuted      = 1108
	ChatMsgSysGroupUserMuted        = 1109
	ChatMsgSysGroupUserCancelMuted  = 1110
	ChatMsgSysGroupAds              = 1111
	ChatMsgSysGroupRenameChange     = 1112
	ChatMsgSysGroupAvatarChange     = 1113
	ChatMsgSysGroupAvatarRoleChange = 1113
)
