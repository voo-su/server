package entity

import "voo.su/internal/constant"

func GetMediaType(ext string) int {
	if val, ok := constant.FileMediaMap[ext]; ok {
		return val
	}

	return constant.MediaFileOther
}
