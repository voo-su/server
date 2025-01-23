package ginutil

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IPWhitelistMiddleware(allowedIPs []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		for _, ip := range allowedIPs {
			if clientIP == ip {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, Response{
			Code:    http.StatusForbidden,
			Message: "Access forbidden",
		})
	}
}
