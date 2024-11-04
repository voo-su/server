package entity

import "voo.su/internal/constant"

func GetMediaType(ext string) int {
	if val, ok := constant.FileMediaMap[ext]; ok {
		return val
	}

	return constant.MediaFileOther
}

var fileSystemDriveMap = map[string]int{
	"local": constant.FileDriveLocal,
}

func FileDriveMode(drive string) int {
	if val, ok := fileSystemDriveMap[drive]; ok {
		return val
	}

	return 0
}
