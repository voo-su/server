package entity

const (
	ChatPrivateMode = 1
	ChatGroupMode   = 2
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
	PushEventContactRequest    = "voo.contact.request"
	PushEventContactStatus     = "voo.contact.status"
	PushEventGroupChatRequest  = "voo.group_chat.request"
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
	ChatMsgSysGroupNotice            = 1111
	ChatMsgSysGroupTransfer          = 1113
)

var ChatMsgTypeMapping = map[int]string{
	ChatMsgTypeImage:                 "Фотография",
	ChatMsgTypeAudio:                 "Аудиозапись",
	ChatMsgTypeVideo:                 "Видео",
	ChatMsgTypeFile:                  "Файл",
	ChatMsgTypeLocation:              "Сообщение с местоположением",
	ChatMsgTypeCard:                  "Сообщение с контактной информацией",
	ChatMsgTypeForward:               "Пересланное сообщение",
	ChatMsgTypeLogin:                 "Уведомление о входе в систему",
	ChatMsgTypeVote:                  "Опрос",
	ChatMsgTypeCode:                  "Сообщение с кодом",
	ChatMsgTypeMixed:                 "Фотографии",
	ChatMsgSysText:                   "Системное сообщение",
	ChatMsgSysGroupCreate:            "Создание группы",
	ChatMsgSysGroupMemberJoin:        "Присоединение к группе",
	ChatMsgSysGroupMemberQuit:        "Выход из группы",
	ChatMsgSysGroupMemberKicked:      "Исключение из группы",
	ChatMsgSysGroupMessageRevoke:     "Отзыв сообщения",
	ChatMsgSysGroupDismissed:         "Распуск группы",
	ChatMsgSysGroupMuted:             "Группа замолчана",
	ChatMsgSysGroupCancelMuted:       "Группа размолчана",
	ChatMsgSysGroupMemberMuted:       "Замолчание участника группы",
	ChatMsgSysGroupMemberCancelMuted: "Размолчание участника группы отменено",
	ChatMsgSysGroupNotice:            "Объявление в группе",
}
