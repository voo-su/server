package entity

import (
	"github.com/google/uuid"
	"time"
	"voo.su/internal/constant"
)

type MessageItem struct {
	Id         int          `json:"id"`
	Sequence   int          `json:"sequence"`
	ChatType   int          `json:"chat_type"`
	MsgType    int          `json:"msg_type"`
	UserId     int          `json:"user_id"`
	ReceiverId int          `json:"receiver_id"`
	Username   string       `json:"username"`
	Name       string       `json:"name"`
	Surname    string       `json:"surname"`
	Avatar     string       `json:"avatar"`
	IsRevoke   int          `json:"is_revoke"`
	IsMark     int          `json:"is_mark"`
	IsRead     int          `json:"is_read"`
	Content    string       `json:"content"`
	CreatedAt  string       `json:"created_at"`
	Reply      *Reply       `json:"reply,omitempty"`
	Media      *MediaItem   `json:"media,omitempty"`
	Service    *ServiceItem `json:"service,omitempty"`
}

type MediaType int

const (
	MediaTypeImage MediaType = constant.ChatMsgTypeImage
	MediaTypeVideo MediaType = constant.ChatMsgTypeVideo
	MediaTypeAudio MediaType = constant.ChatMsgTypeAudio
	MediaTypeFile  MediaType = constant.ChatMsgTypeFile
	MediaTypeVote  MediaType = constant.ChatMsgTypeVote
)

type MediaItem struct {
	Type     MediaType `json:"type"`
	FileId   uuid.UUID `json:"file_id"`
	MimeType string    `json:"mime_type"`
	Url      string    `json:"url"`
	Name     string    `json:"name"`
	Size     int       `json:"size"`
	Duration int       `json:"duration"`
	Width    int32     `json:"width"`
	Height   int32     `json:"height"`
	Cover    string    `json:"cover"`
	Drive    int       `json:"drive"`

	//
	Detail     DetailVote `json:"detail,omitempty"`
	Statistics any        `json:"statistics,omitempty"`
	VoteUsers  []int      `json:"vote_users,omitempty"`
}

type DetailVote struct {
	Id           int    `json:"id"`
	MessageId    int    `json:"message_id"`
	Title        string `json:"title"`
	AnswerMode   int    `json:"answer_mode"`
	Status       int    `json:"status"`
	AnswerOption []any  `json:"answer_option"`
	AnswerNum    int    `json:"answer_num"`
	AnsweredNum  int    `json:"answered_num"`
}

type ServiceType int

const (
	ServiceTypeMsgSysText       ServiceType = constant.ChatMsgSysText
	ServiceTypeChatMsgTypeLogin ServiceType = constant.ChatMsgTypeLogin
)

type ServiceItem struct {
	Type ServiceType `json:"type"`
	//
	Ip      string  `json:"ip,omitempty"`
	Agent   string  `json:"agent,omitempty"`
	Address *string `json:"address,omitempty"`

	CreatedAt time.Time `json:"created_at"`
}

type QueryGetHistoryOpt struct {
	ChatType   int
	UserId     int
	ReceiverId int
	MsgType    []int
	MessageId  int
	Limit      int
}

type QueryMessageItem struct {
	Id         int       `json:"id"`
	Sequence   int       `json:"sequence"`
	ChatType   int       `json:"chat_type"`
	MsgType    int       `json:"msg_type"`
	UserId     int       `json:"user_id"`
	ReceiverId int       `json:"receiver_id"`
	IsRevoke   int       `json:"is_revoke"`
	IsMark     int       `json:"is_mark"`
	IsRead     int       `json:"is_read"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
}

type MessageExtraGroupMembers struct {
	UserId   int    `gorm:"column:user_id;" json:"user_id"`
	Username string `gorm:"column:username;" json:"username"`
}

type MessageReply struct {
	Id           int    `gorm:"id"`
	Content      string `gorm:"content"`
	MsgType      int    `gorm:"msg_type"`
	UserId       int    `gorm:"user_id"`
	NewMessageId int    `gorm:"new_message_id"`
}

type Reply struct {
	Id       int    `json:"id"`
	MsgType  int    `json:"msg_type"`
	Content  string `json:"content"`
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
}

type MessageExtraForward struct {
	ChatType   int              `json:"chat_type"`
	UserId     int              `json:"user_id"`
	ReceiverId int              `json:"receiver_id"`
	Ids        []int            `json:"ids"`
	Messages   []map[string]any `json:"messages"`
}

type MessageExtraFile struct {
	Name   string `json:"name"`
	Drive  int    `json:"drive"`
	Suffix string `json:"suffix"`
	Size   int    `json:"size"`
	Path   string `json:"path"`
}

type MessageExtraImage struct {
	Name   string `json:"name"`
	Suffix string `json:"suffix"`
	Size   int    `json:"size"`
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type MessageExtraGroupCreate struct {
	OwnerId   int                        `json:"owner_id"`
	OwnerName string                     `json:"owner_name"`
	Members   []MessageExtraGroupMembers `json:"members"`
}

type MessageExtraGroupJoin struct {
	OwnerId   int                        `json:"owner_id"`
	OwnerName string                     `json:"owner_name"`
	Members   []MessageExtraGroupMembers `json:"members"`
}

type MessageExtraGroupMuted struct {
	OwnerId   int    `json:"owner_id"`
	OwnerName string `json:"owner_name"`
}

type MessageExtraGroupCancelMuted struct {
	OwnerId   int    `json:"owner_id"`
	OwnerName string `json:"owner_name"`
}

type MessageExtraGroupMemberQuit struct {
	OwnerId   int    `json:"owner_id"`
	OwnerName string `json:"owner_name"`
}

type MessageExtraGroupMemberKicked struct {
	OwnerId   int                        `json:"owner_id"`
	OwnerName string                     `json:"owner_name"`
	Members   []MessageExtraGroupMembers `json:"members"`
}

type MessageExtraGroupAds struct {
	OwnerId   int    `json:"owner_id"`
	OwnerName string `json:"owner_name"`
	Title     string `json:"title"`
	Content   string `json:"content"`
}

type MessageExtraMixedItem struct {
	Type    int    `json:"type"`
	Content string `json:"content"`
	Link    string `json:"link,omitempty"`
}

type MessageExtraMixed struct {
	Items []*MessageExtraMixedItem `json:"items"`
}

type MessageAccess struct {
	ChatType          int
	UserId            int
	ReceiverId        int
	IsVerifyGroupMute bool
}

type MessageReceiver struct {
	ChatType   int32
	ReceiverId int32
}

type SendText struct {
	Receiver     MessageReceiver
	Content      string
	ReplyToMsgId int64
}

type SendImage struct {
	Receiver     MessageReceiver
	Url          string
	Width        int32
	Height       int32
	Content      string
	ReplyToMsgId int64
	FileId       uuid.UUID
}

type SendVideo struct {
	Receiver MessageReceiver
	Url      string
	Duration int32
	Size     int32
	Cover    string
	Content  string
	FileId   uuid.UUID
}

type SendAudio struct {
	Receiver MessageReceiver
	Url      string
	Size     int32
	Content  string
	FileId   uuid.UUID
}

type SendFile struct {
	Receiver MessageReceiver
	UploadId string
}

type SendBotFile struct {
	Receiver     MessageReceiver
	Drive        int
	OriginalName string
	FileExt      string
	FileSize     int
	FilePath     string
	Content      string
	FileId       uuid.UUID
}

type SendLogin struct {
	Ip      string
	Agent   string
	Address *string
}

type Mention struct {
	All  int32
	Uids []int32
}

type TextMessageRequest struct {
	Receiver     *MessageReceiver
	Type         string
	Content      string
	Mention      *Mention
	ReplyToMsgId int64
}
