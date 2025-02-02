package ginutil

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CorsMiddleware(origin string, headers string, methods string, credentials string, maxAge string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Headers", headers)
		c.Header("Access-Control-Allow-Methods", methods)
		c.Header("Access-Control-Allow-Credentials", credentials)
		c.Header("Access-Control-Max-Age", maxAge)
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}
