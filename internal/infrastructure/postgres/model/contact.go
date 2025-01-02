// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package model

import "time"

type Contact struct {
	Id        int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	UserId    int       `gorm:"column:user_id;default:0;NOT NULL" json:"user_id"`
	FriendId  int       `gorm:"column:friend_id;default:0;NOT NULL" json:"friend_id"`
	Remark    string    `gorm:"column:remark;NOT NULL" json:"remark"`
	Status    int       `gorm:"column:status;default:0;NOT NULL" json:"status"`
	FolderId  int       `gorm:"column:group_id;default:0;NOT NULL" json:"group_id"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"updated_at"`
}

func (Contact) TableName() string {
	return "contacts"
}
