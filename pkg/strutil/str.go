package strutil

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func GenValidateCode(length int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	newRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	var sb strings.Builder
	for i := 0; i < length; i++ {
		_, _ = fmt.Fprintf(&sb, "%d", numeric[newRand.Intn(10)])
	}

	return sb.String()
}

func Random(length int) string {
	var result []byte
	bytes := []byte("0123456789abcdefghijklmnopqrstuvwxyz")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}

	return string(result)
}

func GenImageName(ext string, width, height int) string {
	return fmt.Sprintf("%s_%dx%d.%s", uuid.New().String(), width, height, ext)
}

func GenFileName(ext string) string {
	return fmt.Sprintf("%s.%s", uuid.New().String(), ext)
}

func MtSubstr(value string, start, end int) string {
	if start > end {
		return ""
	}
	str := []rune(value)
	if length := len(str); end > length {
		end = length
	}

	return string(str[start:end])
}

func BoolToInt(value bool) int {
	if value {
		return 1
	}

	return 0
}

func FileSuffix(filename string) string {
	return strings.TrimPrefix(path.Ext(filename), ".")
}

func NewMsgId() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

func NewUuid() string {
	return uuid.New().String()
}

func GenMediaObjectName(ext string, width, height int) string {
	var (
		mediaType = "files"
		fileName  = GenFileName(ext)
	)

	switch ext {
	case "png", "jpeg", "jpg", "gif", "webp", "svg", "ico", "bmp", "tiff", "raw", "heif", "heic":
		mediaType = "images"
		fileName = GenImageName(ext, width, height)
	case "mp3", "wav", "aac", "ogg", "flac", "m4a", "opus", "amr", "wma":
		mediaType = "audio"
	case "mp4", "avi", "mov", "wmv", "mkv", "flv", "webm", "3gp", "mpg", "mpeg", "rm", "rmvb":
		mediaType = "videos"
	}

	return fmt.Sprintf("%s/%s/%s", mediaType, time.Now().Format("2006/01/02"), fileName)
}

func ExtractFileExtension(fileName string) string {
	ext := strings.TrimPrefix(filepath.Ext(fileName), ".")
	return strings.ToLower(ext)
}
