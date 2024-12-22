package entity

type DialogRecordExtraGroupMembers struct {
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

type DialogRecordExtraCode struct {
	Lang string `json:"translations"`
	Code string `json:"code"`
}

type DialogRecordExtraLocation struct {
	Longitude   string `json:"longitude"`
	Latitude    string `json:"latitude"`
	Description string `json:"description"`
}

type DialogRecordExtraForward struct {
	DialogType int              `json:"dialog_type"`
	UserId     int              `json:"user_id"`
	ReceiverId int              `json:"receiver_id"`
	MsgIds     []int            `json:"msg_ids"`
	Records    []map[string]any `json:"records"`
}

type DialogRecordExtraLogin struct {
	IP       string `json:"ip"`
	Agent    string `json:"agent"`
	Address  string `json:"address"`
	Datetime string `json:"datetime"`
}

type DialogRecordExtraCard struct {
	UserId int `json:"user_id"`
}

type DialogRecordExtraFile struct {
	Name   string `json:"name"`
	Drive  int    `json:"drive"`
	Suffix string `json:"suffix"`
	Size   int    `json:"size"`
	Path   string `json:"path"`
}

type DialogRecordExtraImage struct {
	Name   string `json:"name"`
	Suffix string `json:"suffix"`
	Size   int    `json:"size"`
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type DialogRecordExtraAudio struct {
	Name     string `json:"name"`
	Suffix   string `json:"suffix"`
	Size     int    `json:"size"`
	Url      string `json:"url"`
	Duration int    `json:"duration"`
}

type DialogRecordExtraVideo struct {
	Name     string `json:"name"`
	Cover    string `json:"cover"`
	Suffix   string `json:"suffix"`
	Size     int    `json:"size"`
	Url      string `json:"url"`
	Duration int    `json:"duration"`
}

type DialogRecordExtraGroupCreate struct {
	OwnerId   int                             `json:"owner_id"`
	OwnerName string                          `json:"owner_name"`
	Members   []DialogRecordExtraGroupMembers `json:"members"`
}

type DialogRecordExtraGroupJoin struct {
	OwnerId   int                             `json:"owner_id"`
	OwnerName string                          `json:"owner_name"`
	Members   []DialogRecordExtraGroupMembers `json:"members"`
}

type DialogRecordExtraGroupTransfer struct {
	OldOwnerId   int    `json:"old_owner_id"`
	OldOwnerName string `json:"old_owner_name"`
	NewOwnerId   int    `json:"new_owner_id"`
	NewOwnerName string `json:"new_owner_name"`
}

type DialogRecordExtraGroupMuted struct {
	OwnerId   int    `json:"owner_id"`
	OwnerName string `json:"owner_name"`
}

type DialogRecordExtraGroupCancelMuted struct {
	OwnerId   int    `json:"owner_id"`
	OwnerName string `json:"owner_name"`
}

type DialogRecordExtraGroupMemberMuted struct {
	OwnerId   int                             `json:"owner_id"`
	OwnerName string                          `json:"owner_name"`
	Members   []DialogRecordExtraGroupMembers `json:"members"`
}

type DialogRecordExtraGroupMemberCancelMuted struct {
	OwnerId   int                             `json:"owner_id"`
	OwnerName string                          `json:"owner_name"`
	Members   []DialogRecordExtraGroupMembers `json:"members"`
}

type DialogRecordExtraGroupDismissed struct {
	OwnerId   int    `json:"owner_id"`
	OwnerName string `json:"owner_name"`
}

type DialogRecordExtraGroupMemberQuit struct {
	OwnerId   int    `json:"owner_id"`
	OwnerName string `json:"owner_name"`
}

type DialogRecordExtraGroupMemberKicked struct {
	OwnerId   int                             `json:"owner_id"`
	OwnerName string                          `json:"owner_name"`
	Members   []DialogRecordExtraGroupMembers `json:"members"`
}

type DialogRecordExtraGroupMessageRevoke struct {
	OwnerId         int    `json:"owner_id"`
	OwnerName       string `json:"owner_name"`
	RevokeMessageId string `json:"revoke_message_id"`
}

type DialogRecordExtraGroupAds struct {
	OwnerId   int    `json:"owner_id"`
	OwnerName string `json:"owner_name"`
	Title     string `json:"title"`
	Content   string `json:"content"`
}

type DialogRecordExtraText struct {
	Content string `json:"content"`
}

type DialogRecordExtraMixedItem struct {
	Type    int    `json:"type"`
	Content string `json:"content"`
	Link    string `json:"link,omitempty"`
}

type DialogRecordExtraMixed struct {
	Items []*DialogRecordExtraMixedItem `json:"items"`
}
