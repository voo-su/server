package jwtutil

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
	"voo.su/pkg/locale"
)

const JWTSession = "__JWT_SESSION__"

type IStorage interface {
	IsBlackList(ctx context.Context, token string) bool
}

type JSession struct {
	Uid       int    `json:"uid"`
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
}

type Options jwt.RegisteredClaims

type AuthClaims struct {
	Guard string `json:"guard"`
	jwt.RegisteredClaims
}

func NewNumericDate(t time.Time) *jwt.NumericDate {
	return jwt.NewNumericDate(t)
}

func GenerateToken(guard string, secret string, ops *Options) string {
	claims := AuthClaims{
		Guard: guard,
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  ops.Audience,
			ExpiresAt: ops.ExpiresAt,
			ID:        ops.ID,
			IssuedAt:  ops.IssuedAt,
			Issuer:    ops.Issuer,
			NotBefore: ops.NotBefore,
			Subject:   ops.Subject,
		},
	}

	tokenString, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))

	return tokenString
}

func ParseToken(token string, secret string) (*AuthClaims, error) {
	data, err := jwt.ParseWithClaims(token, &AuthClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("неожиданный метод подписи: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if claims, ok := data.Claims.(*AuthClaims); ok && data.Valid {
		return claims, nil
	}

	return nil, err
}

func Verify(locale locale.ILocale, guard string, secret string, token string) (*AuthClaims, error) {
	if token == "" {
		return nil, errors.New(locale.Localize("authorization_required"))
	}

	claims, err := ParseToken(token, secret)
	if err != nil {
		return nil, err
	}

	if claims.Guard != guard || claims.Valid() != nil {
		return nil, errors.New(locale.Localize("authorization_required"))
	}

	return claims, nil
}
