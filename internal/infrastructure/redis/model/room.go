package model

import "voo.su/internal/constant"

type RoomOption struct {
	Channel  string
	RoomType constant.RoomType
	Number   string
	Sid      string
	Cid      int64
}
