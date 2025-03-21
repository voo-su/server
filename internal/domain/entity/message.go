package entity

import "time"

type MessageItem struct {
	Id         int    `json:"id"`
	Sequence   int    `json:"sequence"`
	MsgId      string `json:"msg_id"`
	ChatType   int    `json:"chat_type"`
	MsgType    int    `json:"msg_type"`
	UserId     int    `json:"user_id"`
	ReceiverId int    `json:"receiver_id"`
	Username   string `json:"username"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Avatar     string `json:"avatar"`
	IsRevoke   int    `json:"is_revoke"`
	IsMark     int    `json:"is_mark"`
	IsRead     int    `json:"is_read"`
	Content    string `json:"content"`
	CreatedAt  string `json:"created_at"`
	Extra      any    `json:"extra"`
	Extra0     string `json:"extra0"`
	QuoteId    string `json:"quote_id"`
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
	MsgId      string    `json:"msg_id"`
	Sequence   int64     `json:"sequence"`
	ChatType   int       `json:"chat_type"`
	MsgType    int       `json:"msg_type"`
	UserId     int       `json:"user_id"`
	ReceiverId int       `json:"receiver_id"`
	IsRevoke   int       `json:"is_revoke"`
	IsMark     int       `json:"is_mark"`
	IsRead     int       `json:"is_read"`
	QuoteId    string    `json:"quote_id"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
	Username   string    `json:"username"`
	Name       string    `json:"name"`
	Surname    string    `json:"surname"`
	Avatar     string    `json:"avatar"`
	Extra      string    `json:"extra"`
}

type MessageExtraGroupMembers struct {
	UserId   int    `gorm:"column:user_id;" json:"user_id"`
	Username string `gorm:"column:username;" json:"username"`
}

type Reply struct {
	UserId   int    `json:"user_id,omitempty"`
	Username string `json:"username,omitempty"`
	MsgType  int    `json:"msg_type,omitempty"`
	Content  string `json:"content,omitempty"`
	MsgId    string `json:"msg_id,omitempty"`
}

type MessageExtraCode struct {
	Lang string `json:"translations"`
	Code string `json:"code"`
}

type MessageExtraLocation struct {
	Longitude   string `json:"longitude"`
	Latitude    string `json:"latitude"`
	Description string `json:"description"`
}

type MessageExtraForward struct {
	ChatType   int              `json:"chat_type"`
	UserId     int              `json:"user_id"`
	ReceiverId int              `json:"receiver_id"`
	MsgIds     []int            `json:"msg_ids"`
	Messages   []map[string]any `json:"messages"`
}

type MessageExtraLogin struct {
	IP       string `json:"ip"`
	Agent    string `json:"agent"`
	Address  string `json:"address"`
	Datetime string `json:"datetime"`
}

type MessageExtraCard struct {
	UserId int `json:"user_id"`
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

type MessageExtraAudio struct {
	Name     string `json:"name"`
	Suffix   string `json:"suffix"`
	Size     int    `json:"size"`
	Url      string `json:"url"`
	Duration int    `json:"duration"`
}

type MessageExtraVideo struct {
	Name     string `json:"name"`
	Cover    string `json:"cover"`
	Suffix   string `json:"suffix"`
	Size     int    `json:"size"`
	Url      string `json:"url"`
	Duration int    `json:"duration"`
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

type MessageExtraGroupTransfer struct {
	OldOwnerId   int    `json:"old_owner_id"`
	OldOwnerName string `json:"old_owner_name"`
	NewOwnerId   int    `json:"new_owner_id"`
	NewOwnerName string `json:"new_owner_name"`
}

type MessageExtraGroupMuted struct {
	OwnerId   int    `json:"owner_id"`
	OwnerName string `json:"owner_name"`
}

type MessageExtraGroupCancelMuted struct {
	OwnerId   int    `json:"owner_id"`
	OwnerName string `json:"owner_name"`
}

type MessageExtraGroupMemberMuted struct {
	OwnerId   int                        `json:"owner_id"`
	OwnerName string                     `json:"owner_name"`
	Members   []MessageExtraGroupMembers `json:"members"`
}

type MessageExtraGroupMemberCancelMuted struct {
	OwnerId   int                        `json:"owner_id"`
	OwnerName string                     `json:"owner_name"`
	Members   []MessageExtraGroupMembers `json:"members"`
}

type MessageExtraGroupDismissed struct {
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

type MessageExtraGroupMessageRevoke struct {
	OwnerId         int    `json:"owner_id"`
	OwnerName       string `json:"owner_name"`
	RevokeMessageId string `json:"revoke_message_id"`
}

type MessageExtraGroupAds struct {
	OwnerId   int    `json:"owner_id"`
	OwnerName string `json:"owner_name"`
	Title     string `json:"title"`
	Content   string `json:"content"`
}

type MessageExtraText struct {
	Content string `json:"content"`
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
	QuoteId      string
	ReplyToMsgId int64
}

type SendImage struct {
	Receiver     MessageReceiver
	Url          string
	Width        int32
	Height       int32
	QuoteId      string
	Content      string
	ReplyToMsgId int64
}

type SendVideo struct {
	Receiver MessageReceiver
	Url      string
	Duration int32
	Size     int32
	Cover    string
	Content  string
}

type SendAudio struct {
	Receiver MessageReceiver
	Url      string
	Size     int32
	Content  string
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
}

type SendLogin struct {
	Ip      string
	Agent   string
	Address string
}

type Mention struct {
	All  int32
	Uids []int32
}

type TextMessageRequest struct {
	Receiver *MessageReceiver
	Type     string
	Content  string
	Mention  *Mention
	QuoteId  string
}
