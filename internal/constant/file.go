package constant

const (
	MediaFileImage = 1
	MediaFileVideo = 2
	MediaFileAudio = 3
	MediaFileOther = 4
)

var FileMediaMap = map[string]int{
	"gif":  MediaFileImage,
	"jpg":  MediaFileImage,
	"jpeg": MediaFileImage,
	"png":  MediaFileImage,
	"webp": MediaFileImage,
	"mp3":  MediaFileAudio,
	"wav":  MediaFileAudio,
	"mp4":  MediaFileVideo,
}
