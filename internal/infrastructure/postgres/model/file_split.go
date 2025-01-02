// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package model

import "time"

type FileSplit struct {
	Id           int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Type         int       `gorm:"column:type;default:1;NOT NULL" json:"type"`
	Drive        int       `gorm:"column:drive;default:1;NOT NULL" json:"drive"`
	UploadId     string    `gorm:"column:upload_id;NOT NULL" json:"upload_id"`
	UserId       int       `gorm:"column:user_id;default:0;NOT NULL" json:"user_id"`
	OriginalName string    `gorm:"column:original_name;NOT NULL" json:"original_name"`
	SplitIndex   int       `gorm:"column:split_index;default:0;NOT NULL" json:"split_index"`
	SplitNum     int       `gorm:"column:split_num;default:0;NOT NULL" json:"split_num"`
	Path         string    `gorm:"column:path;NOT NULL" json:"path"`
	FileExt      string    `gorm:"column:file_ext;NOT NULL" json:"file_ext"`
	FileSize     int64     `gorm:"column:file_size;NOT NULL" json:"file_size"`
	IsDelete     int       `gorm:"column:is_delete;default:0;NOT NULL" json:"is_delete"`
	Attr         string    `gorm:"column:attr;NOT NULL" json:"attr"`
	CreatedAt    time.Time `gorm:"column:created_at;NOT NULL" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;NOT NULL" json:"updated_at"`
}

func (FileSplit) TableName() string {
	return "file_splits"
}
