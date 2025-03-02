package model

import "time"

type FileSplit struct {
	Id           int       `gorm:"primaryKey"`
	Type         int       `gorm:"column:type;DEFAULT:1;NOT NULL"`
	Drive        int       `gorm:"column:drive;DEFAULT:1;NOT NULL"`
	UploadId     string    `gorm:"column:upload_id;NOT NULL"`
	UserId       int       `gorm:"column:user_id;DEFAULT:0;NOT NULL"`
	OriginalName string    `gorm:"column:original_name;NOT NULL"`
	SplitIndex   int       `gorm:"column:split_index;DEFAULT:0;NOT NULL"`
	SplitNum     int       `gorm:"column:split_num;DEFAULT:0;NOT NULL"`
	Path         string    `gorm:"column:path;NOT NULL"`
	FileExt      string    `gorm:"column:file_ext;NOT NULL"`
	FileSize     int64     `gorm:"column:file_size;NOT NULL"`
	IsDelete     int       `gorm:"column:is_delete;DEFAULT:0;NOT NULL"`
	Attr         string    `gorm:"column:attr;NOT NULL"`
	CreatedAt    time.Time `gorm:"column:created_at;NOT NULL"`
	UpdatedAt    time.Time `gorm:"column:updated_at;NOT NULL"`
}

func (FileSplit) TableName() string {
	return "file_splits"
}
