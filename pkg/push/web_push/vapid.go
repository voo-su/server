package web_push

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"math/big"
	"net/url"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// GenerateVAPIDKeys создаёт пару ключей VAPID: приватный и публичный
func GenerateVAPIDKeys() (privateKey, publicKey string, err error) {
	curve := elliptic.P256()
	private, x, y, err := elliptic.GenerateKey(curve, rand.Reader)
	if err != nil {
		return
	}

	public := elliptic.Marshal(curve, x, y)
	publicKey = base64.RawURLEncoding.EncodeToString(public)
	privateKey = base64.RawURLEncoding.EncodeToString(private)

	return
}

// Генерирует публичные и приватные ключи ECDSA для шифрования JWT
func generateVAPIDHeaderKeys(privateKey []byte) *ecdsa.PrivateKey {
	// Публичный ключ
	curve := elliptic.P256()
	px, py := curve.ScalarMult(
		curve.Params().Gx,
		curve.Params().Gy,
		privateKey,
	)

	pubKey := ecdsa.PublicKey{
		Curve: curve,
		X:     px,
		Y:     py,
	}

	// Приватный ключ
	d := &big.Int{}
	d.SetBytes(privateKey)

	return &ecdsa.PrivateKey{
		PublicKey: pubKey,
		D:         d,
	}
}

// Создаёт заголовок авторизации VAPID
func getVAPIDAuthorizationHeader(
	endpoint,
	subscriber,
	vapidPublicKey,
	vapidPrivateKey string,
	expiration time.Time,
) (string, error) {
	// Создаём JWT токен
	subURL, err := url.Parse(endpoint)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"aud": fmt.Sprintf("%s://%s", subURL.Scheme, subURL.Host),
		"exp": expiration.Unix(),
		"sub": fmt.Sprintf("mailto:%s", subscriber),
	})

	// Декодируем приватный ключ VAPID
	decodedVapidPrivateKey, err := decodeVapidKey(vapidPrivateKey)
	if err != nil {
		return "", err
	}

	privateKey := generateVAPIDHeaderKeys(decodedVapidPrivateKey)

	// Подписываем токен приватным ключом
	jwtString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	// Декодируем публичный ключ VAPID
	pubKey, err := decodeVapidKey(vapidPublicKey)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"vapid t=%s, k=%s",
		jwtString,
		base64.RawURLEncoding.EncodeToString(pubKey),
	), nil
}

// Декодирует ключ VAPID из строки base64
func decodeVapidKey(key string) ([]byte, error) {
	bytes, err := base64.URLEncoding.DecodeString(key)
	if err == nil {
		return bytes, nil
	}

	return base64.RawURLEncoding.DecodeString(key)
}
