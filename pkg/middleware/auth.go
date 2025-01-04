package middleware

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
	"net/http"
	"strconv"
	"strings"
	"voo.su/pkg/jwt"
	"voo.su/pkg/locale"
	"voo.su/pkg/response"
)

const JWTSessionConst = "__JWT_SESSION__"

type IStorage interface {
	IsBlackList(ctx context.Context, token string) bool
}

type JSession struct {
	Uid       int    `json:"uid"`
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
}

func Auth(locale locale.ILocale, guard string, secret string, storage IStorage) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := HttpToken(c)
		claims, err := verify(locale, guard, secret, token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Response{
				Code:    http.StatusUnauthorized,
				Message: err.Error(),
			})
			return
		}

		if storage.IsBlackList(c.Request.Context(), token) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Response{
				Code:    http.StatusUnauthorized,
				Message: locale.Localize("authorization_required"),
			})
			return
		}

		uid, err := strconv.Atoi(claims.ID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, response.Response{
				Code:    http.StatusInternalServerError,
				Message: "Ошибка разбора jwt",
			})
			return
		}

		c.Set(JWTSessionConst, &JSession{
			Uid:       uid,
			Token:     token,
			ExpiresAt: claims.ExpiresAt.Unix(),
		})

		c.Next()
	}
}

func HttpToken(c *gin.Context) string {
	token := c.GetHeader("Authorization")
	token = strings.TrimSpace(strings.TrimPrefix(token, "Bearer"))
	if token == "" {
		token = c.DefaultQuery("token", "")
	}

	return token
}

func GrpcToken(ctx context.Context, locale locale.ILocale, guard string, secret string) (*jwt.AuthClaims, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New(locale.Localize("failed_to_fetch_metadata"))
	}

	token := md.Get("Authorization")
	userClaims, err := verify(locale, guard, secret, token[len(token)-1])
	if err != nil {
		return nil, err
	}

	return userClaims, nil
}

func verify(locale locale.ILocale, guard string, secret string, token string) (*jwt.AuthClaims, error) {
	if token == "" {
		return nil, errors.New(locale.Localize("authorization_required"))
	}

	claims, err := jwt.ParseToken(token, secret)
	if err != nil {
		return nil, err
	}

	if claims.Guard != guard || claims.Valid() != nil {
		return nil, errors.New(locale.Localize("authorization_required"))
	}

	return claims, nil
}
