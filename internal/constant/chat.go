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
	DialogRecordDialogTypePrivate = 1
	DialogRecordDialogTypeGroup   = 2
)

const (
	VoteAnswerModeSingleChoice = 0
)

const (
	ChatMsgTypeText                  = 1
	ChatMsgTypeCode                  = 2
	ChatMsgTypeImage                 = 3
	ChatMsgTypeAudio                 = 4
	ChatMsgTypeVideo                 = 5
	ChatMsgTypeFile                  = 6
	ChatMsgTypeLocation              = 7
	ChatMsgTypeCard                  = 8
	ChatMsgTypeForward               = 9
	ChatMsgTypeLogin                 = 10
	ChatMsgTypeVote                  = 11
	ChatMsgTypeMixed                 = 12
	ChatMsgSysText                   = 1000
	ChatMsgSysGroupCreate            = 1101
	ChatMsgSysGroupMemberJoin        = 1102
	ChatMsgSysGroupMemberQuit        = 1103
	ChatMsgSysGroupMemberKicked      = 1104
	ChatMsgSysGroupMessageRevoke     = 1105
	ChatMsgSysGroupDismissed         = 1106
	ChatMsgSysGroupMuted             = 1107
	ChatMsgSysGroupCancelMuted       = 1108
	ChatMsgSysGroupMemberMuted       = 1109
	ChatMsgSysGroupMemberCancelMuted = 1110
	ChatMsgSysGroupAds               = 1111
	ChatMsgSysGroupTransfer          = 1113
)

var ChatMsgTypeMapping = map[int]string{
	ChatMsgTypeImage:                 "Фотография",
	ChatMsgTypeAudio:                 "Аудиозапись",
	ChatMsgTypeVideo:                 "Видео",
	ChatMsgTypeFile:                  "Файл",
	ChatMsgTypeLocation:              "Местоположение",
	ChatMsgTypeCard:                  "Контактная информация",
	ChatMsgTypeForward:               "Пересланное сообщение",
	ChatMsgTypeLogin:                 "Уведомление о входе в систему",
	ChatMsgTypeVote:                  "Опрос",
	ChatMsgTypeCode:                  "Код",
	ChatMsgTypeMixed:                 "Фотографии",
	ChatMsgSysText:                   "Системное сообщение",
	ChatMsgSysGroupCreate:            "Создание группы",
	ChatMsgSysGroupMemberJoin:        "Присоединение к группе",
	ChatMsgSysGroupMemberQuit:        "Выход из группы",
	ChatMsgSysGroupMemberKicked:      "Исключение из группы",
	ChatMsgSysGroupMessageRevoke:     "Отзыв сообщения",
	ChatMsgSysGroupDismissed:         "Удаление группы",
	ChatMsgSysGroupMuted:             "Отключение уведомлений в группе",
	ChatMsgSysGroupCancelMuted:       "Включение уведомлений в группе",
	ChatMsgSysGroupMemberMuted:       "Отключение уведомлений для участника группы",
	ChatMsgSysGroupMemberCancelMuted: "Включение уведомлений для участника группы",
	ChatMsgSysGroupAds:               "Объявление в группе",
}
