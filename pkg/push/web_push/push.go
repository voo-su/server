package web_push

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/hkdf"
)

const MaxMessageSize uint32 = 4096

// saltFunc генерирует соль размером 16 байт
var saltFunc = func() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return salt, err
	}

	return salt, nil
}

type IHTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type Options struct {
	HTTPClient      IHTTPClient
	MessageSize     uint32    // Ограничение размера записи
	Subscriber      string    // Указывается в VAPID JWT токене
	Topic           string    // Устанавливает заголовок Topic для группировки сообщений (опционально)
	TTL             int       // Устанавливает TTL в POST-запросе к конечной точке
	Urgency         Urgency   // Устанавливает заголовок Urgency для изменения приоритета сообщения (опционально)
	VAPIDPublicKey  string    // Публичный ключ VAPID, указывается в заголовке авторизации VAPID
	VAPIDPrivateKey string    // Приватный ключ VAPID, используется для подписи VAPID JWT токена
	VapidExpiration time.Time // Опциональное время истечения VAPID JWT токена (по умолчанию - сейчас + 12 часов)
}

// Keys - это base64-кодированные значения из PushSubscription.getKey()
type Keys struct {
	Auth   string `json:"auth"`
	P256dh string `json:"p256dh"`
}

// Subscription представляет объект PushSubscription из Push API
type Subscription struct {
	Endpoint string `json:"endpoint"`
	Keys     Keys   `json:"keys"`
}

// SendNotification вызывает SendNotificationWithContext с контекстом по умолчанию для обратной совместимости
func SendNotification(message []byte, s *Subscription, options *Options) (*http.Response, error) {
	return SendNotificationWithContext(context.Background(), message, s, options)
}

// SendNotificationWithContext отправляет push-уведомление на конечную точку подписки
// Шифрование сообщения для Web Push и VAPID протоколов.
// ДОПОЛНИТЕЛЬНАЯ ИНФОРМАЦИЯ В RFC8291: https://datatracker.ietf.org/doc/rfc8291
func SendNotificationWithContext(ctx context.Context, message []byte, s *Subscription, options *Options) (*http.Response, error) {
	// Секрет аутентификации (auth_secret)
	authSecret, err := decodeSubscriptionKey(s.Keys.Auth)
	if err != nil {
		return nil, err
	}

	// dh (Diffie Hellman)
	dh, err := decodeSubscriptionKey(s.Keys.P256dh)
	if err != nil {
		return nil, err
	}

	// Генерируем соль размером 16 байт
	salt, err := saltFunc()
	if err != nil {
		return nil, err
	}

	// Создаем общую пару ключей ecdh_secret
	curve := elliptic.P256()

	// Ключи сервера приложения (одноразовые)
	localPrivateKey, x, y, err := elliptic.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, err
	}

	localPublicKey := elliptic.Marshal(curve, x, y)

	// Объединяем ключи приложения с EC публичным ключом получателя
	sharedX, sharedY := elliptic.Unmarshal(curve, dh)
	if sharedX == nil {
		return nil, errors.New("Ошибка Unmarshal: Публичный ключ не является допустимой точкой на кривой")
	}

	// Извлекаем общую секретную часть ECDH
	sx, sy := curve.ScalarMult(sharedX, sharedY, localPrivateKey)
	if !curve.IsOnCurve(sx, sy) {
		return nil, errors.New("Ошибка шифрования: Общий секретный ключ ECDH не находится на кривой")
	}

	mlen := curve.Params().BitSize / 8
	sharedECDHSecret := make([]byte, mlen)
	sx.FillBytes(sharedECDHSecret)

	hash := sha256.New

	// ikm
	prkInfoBuf := bytes.NewBuffer([]byte("WebPush: info\x00"))
	prkInfoBuf.Write(dh)
	prkInfoBuf.Write(localPublicKey)

	prkHKDF := hkdf.New(hash, sharedECDHSecret, authSecret, prkInfoBuf.Bytes())
	ikm, err := getHKDFKey(prkHKDF, 32)
	if err != nil {
		return nil, err
	}

	// Извлечение ключа шифрования контента
	contentEncryptionKeyInfo := []byte("Content-Encoding: aes128gcm\x00")
	contentHKDF := hkdf.New(hash, ikm, salt, contentEncryptionKeyInfo)
	contentEncryptionKey, err := getHKDFKey(contentHKDF, 16)
	if err != nil {
		return nil, err
	}

	// Извлечение Nonce
	nonceInfo := []byte("Content-Encoding: nonce\x00")
	nonceHKDF := hkdf.New(hash, ikm, salt, nonceInfo)
	nonce, err := getHKDFKey(nonceHKDF, 12)
	if err != nil {
		return nil, err
	}

	// Шифрование
	c, err := aes.NewCipher(contentEncryptionKey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	// Получаем размер записи
	messageSize := options.MessageSize
	if messageSize == 0 {
		messageSize = MaxMessageSize
	}

	recordLength := int(messageSize) - 16

	// Заголовок шифрования контента
	recordBuf := bytes.NewBuffer(salt)

	rs := make([]byte, 4)
	binary.BigEndian.PutUint32(rs, messageSize)

	recordBuf.Write(rs)
	recordBuf.Write([]byte{byte(len(localPublicKey))})
	recordBuf.Write(localPublicKey)

	dataBuf := bytes.NewBuffer(message)

	// Заполняем контент до максимального размера записи - 16 - заголовок
	// Завершающий разделитель
	dataBuf.Write([]byte("\x02"))
	if err := pad(dataBuf, recordLength-recordBuf.Len()); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal([]byte{}, nonce, dataBuf.Bytes(), nil)
	recordBuf.Write(ciphertext)

	req, err := http.NewRequest("POST", s.Endpoint, recordBuf)
	if err != nil {
		return nil, err
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	req.Header.Set("Content-Encoding", "aes128gcm")
	req.Header.Set("Content-Length", strconv.Itoa(len(ciphertext)))
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("TTL", strconv.Itoa(options.TTL))

	if len(options.Topic) > 0 {
		req.Header.Set("Topic", options.Topic)
	}

	if isValidUrgency(options.Urgency) {
		req.Header.Set("Urgency", string(options.Urgency))
	}

	expiration := options.VapidExpiration
	if expiration.IsZero() {
		expiration = time.Now().Add(time.Hour * 12)
	}

	vapidAuthHeader, err := getVAPIDAuthorizationHeader(
		s.Endpoint,
		options.Subscriber,
		options.VAPIDPublicKey,
		options.VAPIDPrivateKey,
		expiration,
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", vapidAuthHeader)

	var client IHTTPClient
	if options.HTTPClient != nil {
		client = options.HTTPClient
	} else {
		client = &http.Client{}
	}

	return client.Do(req)
}

// decodeSubscriptionKey декодирует base64 ключ подписки.
// При необходимости добавляет "=" для URL-декодирования
func decodeSubscriptionKey(key string) ([]byte, error) {
	// "=" padding
	buf := bytes.NewBufferString(key)
	if rem := len(key) % 4; rem != 0 {
		buf.WriteString(strings.Repeat("=", 4-rem))
	}

	_bytes, err := base64.StdEncoding.DecodeString(buf.String())
	if err == nil {
		return _bytes, nil
	}

	return base64.URLEncoding.DecodeString(buf.String())
}

// Возвращает ключ длиной "length" из функции hkdf
func getHKDFKey(hkdf io.Reader, length int) ([]byte, error) {
	key := make([]byte, length)
	n, err := io.ReadFull(hkdf, key)
	if n != len(key) || err != nil {
		return key, err
	}

	return key, nil
}

// Дополняет буфер payload до указанной длины
func pad(payload *bytes.Buffer, maxPadLen int) error {
	payloadLen := payload.Len()
	if payloadLen > maxPadLen {
		return errors.New("размер данных превышает максимально допустимую длину")
	}

	padLen := maxPadLen - payloadLen

	padding := make([]byte, padLen)
	payload.Write(padding)

	return nil
}
