package model

import "time"

const (
	GroupMemberQuitStatusYes = 1
	GroupMemberQuitStatusNo  = 0
	GroupMemberMuteStatusYes = 1
	GroupMemberMuteStatusNo  = 0
)

type GroupChatMember struct {
	Id      int `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	GroupId int `gorm:"column:group_id;default:0;NOT NULL" json:"group_id"`
	UserId  int `gorm:"column:user_id;default:0;NOT NULL" json:"user_id"`
	Leader  int `gorm:"column:leader;default:0;NOT NULL" json:"leader"`
	//UserCard    string    `gorm:"column:user_card;NOT NULL" json:"user_card"`
	IsQuit      int       `gorm:"column:is_quit;default:0;NOT NULL" json:"is_quit"`
	IsMute      int       `gorm:"column:is_mute;default:0;NOT NULL" json:"is_mute"`
	MinRecordId int       `gorm:"column:min_record_id;default:0;NOT NULL" json:"min_record_id"`
	JoinTime    time.Time `gorm:"column:join_time;" json:"join_time"`
	CreatedAt   time.Time `gorm:"column:created_at;NOT NULL" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;NOT NULL" json:"updated_at"`
}

func (GroupChatMember) TableName() string {
	return "group_chat_members"
}

type MemberItem struct {
	Id       string `json:"id"`
	UserId   int    `json:"user_id"`
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
	Gender   int    `json:"gender"`
	About    string `json:"about"`
	Leader   int    `json:"leader"`
	IsMute   int    `json:"is_mute"`
	//UserCard string `json:"user_card"`
}
