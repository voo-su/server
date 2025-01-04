package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"voo.su/pkg/ginutil"
	"voo.su/pkg/jwtutil"
	"voo.su/pkg/locale"
)

func Auth(locale locale.ILocale, guard string, secret string, storage jwtutil.IStorage) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		token = strings.TrimSpace(strings.TrimPrefix(token, "Bearer"))

		claims, err := jwtutil.Verify(locale, guard, secret, token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, ginutil.Response{
				Code:    http.StatusUnauthorized,
				Message: err.Error(),
			})
			return
		}

		if storage.IsBlackList(c.Request.Context(), token) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, ginutil.Response{
				Code:    http.StatusUnauthorized,
				Message: locale.Localize("authorization_required"),
			})
			return
		}

		uid, err := strconv.Atoi(claims.ID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, ginutil.Response{
				Code:    http.StatusInternalServerError,
				Message: locale.Localize("general_error"),
			})
			return
		}

		c.Set(jwtutil.JWTSession, &jwtutil.JSession{
			Uid:       uid,
			Token:     token,
			ExpiresAt: claims.ExpiresAt.Unix(),
		})

		c.Next()
	}
}
