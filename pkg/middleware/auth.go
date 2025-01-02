// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package middleware

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"voo.su/pkg/jwt"
	"voo.su/pkg/response"
)

const JWTSessionConst = "__JWT_SESSION__"

var (
	ErrorNoLogin = errors.New("Пожалуйста, войдите в систему перед выполнением операции")
)

type IStorage interface {
	IsBlackList(ctx context.Context, token string) bool
}

type JSession struct {
	Uid       int    `json:"uid"`
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
}

func Auth(secret string, guard string, storage IStorage) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := AuthHeaderToken(c)
		claims, err := verify(guard, secret, token)
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
				Message: "Пожалуйста, войдите в систему и попробуйте снова",
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

func AuthHeaderToken(c *gin.Context) string {
	token := c.GetHeader("Authorization")
	token = strings.TrimSpace(strings.TrimPrefix(token, "Bearer"))
	if token == "" {
		token = c.DefaultQuery("token", "")
	}

	return token
}

func verify(guard string, secret string, token string) (*jwt.AuthClaims, error) {
	if token == "" {
		return nil, ErrorNoLogin
	}

	claims, err := jwt.ParseToken(token, secret)
	if err != nil {
		return nil, err
	}

	if claims.Guard != guard || claims.Valid() != nil {
		return nil, ErrorNoLogin
	}

	return claims, nil
}
