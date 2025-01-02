// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

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
