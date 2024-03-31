package entity

const (
	MediaFileImage = 1
	MediaFileVideo = 2
	MediaFileAudio = 3
	MediaFileOther = 4
)

var fileMediaMap = map[string]int{
	"gif":  MediaFileImage,
	"jpg":  MediaFileImage,
	"jpeg": MediaFileImage,
	"png":  MediaFileImage,
	"webp": MediaFileImage,
	"mp3":  MediaFileAudio,
	"wav":  MediaFileAudio,
	"mp4":  MediaFileVideo,
}

func GetMediaType(ext string) int {
	if val, ok := fileMediaMap[ext]; ok {
		return val
	}

	return MediaFileOther
}

const (
	FileDriveLocal = 1
)

var fileSystemDriveMap = map[string]int{
	"local": FileDriveLocal,
}

func FileDriveMode(drive string) int {
	if val, ok := fileSystemDriveMap[drive]; ok {
		return val
	}

	return 0
}
