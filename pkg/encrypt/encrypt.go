package encrypt

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"log"
)

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))

	return hex.EncodeToString(h.Sum(nil))
}

// GenerateToken создает случайный токен длиной 32 байта (256 бит)
func GenerateToken() string {
	token := make([]byte, 32)

	_, err := rand.Read(token)
	if err != nil {
		log.Fatalf("Не удалось сгенерировать токен: %v", err)
	}
	return hex.EncodeToString(token)
}
