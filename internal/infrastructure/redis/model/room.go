// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package model

import "voo.su/internal/constant"

type RoomOption struct {
	Channel  string
	RoomType constant.RoomType
	Number   string
	Sid      string
	Cid      int64
}
