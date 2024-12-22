package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc/metadata"
	"time"
	"voo.su/internal/config"
	"voo.su/internal/infrastructure/postgres/model"
	"voo.su/pkg/locale"
)

type TokenMiddleware struct {
	Conf   *config.Config
	Locale locale.ILocale
}

func NewTokenMiddleware(
	conf *config.Config,
	locale locale.ILocale,
) *TokenMiddleware {
	return &TokenMiddleware{
		Conf:   conf,
		Locale: locale,
	}
}

type UserInfo struct {
	Username string `json:"username"`
	Id       int    `json:"id"`
}

type UserClaims struct {
	*jwt.StandardClaims
	UserInfo
}

func (t *TokenMiddleware) CreateToken(user *model.User) (string, error) {
	claims := &UserClaims{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    "voo.su",
		},
		UserInfo{user.Username, user.Id},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(t.Conf.App.Jwt.Secret))
}

func (t *TokenMiddleware) ParseToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf(t.Locale.Localize("unexpected_signature_method"), token.Header["alg"])
		}

		return []byte(t.Conf.App.Jwt.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

func (t *TokenMiddleware) ValidateToken(ctx context.Context) (*UserClaims, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New(t.Locale.Localize("failed_to_fetch_metadata"))
	}

	token := md.Get("Authorization")
	userClaims, err := t.ParseToken(token[len(token)-1])
	if err != nil {
		return nil, err
	}

	return userClaims, nil
}
