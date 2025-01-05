package model

import "time"

type GroupChatMember struct {
	Id          int       `gorm:"primaryKey"`
	GroupId     int       `gorm:"column:group_id;default:0;NOT NULL"`
	UserId      int       `gorm:"column:user_id;default:0;NOT NULL"`
	Leader      int       `gorm:"column:leader;default:0;NOT NULL"`
	IsQuit      int       `gorm:"column:is_quit;default:0;NOT NULL"`
	IsMute      int       `gorm:"column:is_mute;default:0;NOT NULL"`
	MinRecordId int       `gorm:"column:min_record_id;default:0;NOT NULL"`
	JoinTime    time.Time `gorm:"column:join_time;"`
	CreatedAt   time.Time `gorm:"column:created_at;NOT NULL"`
	UpdatedAt   time.Time `gorm:"column:updated_at;NOT NULL"`
}

func (GroupChatMember) TableName() string {
	return "group_chat_members"
}
