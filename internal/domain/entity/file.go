// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package entity

import "voo.su/internal/constant"

func GetMediaType(ext string) int {
	if val, ok := constant.FileMediaMap[ext]; ok {
		return val
	}

	return constant.MediaFileOther
}
