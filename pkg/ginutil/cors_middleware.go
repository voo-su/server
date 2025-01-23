package ginutil

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ICorsOptions interface {
	GetOrigin() string

	GetHeaders() string

	GetMethods() string

	GetCredentials() string

	GetMaxAge() string
}

func CorsMiddleware(options ICorsOptions) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", options.GetOrigin())
		c.Header("Access-Control-Allow-Headers", options.GetHeaders())
		c.Header("Access-Control-Allow-Methods", options.GetMethods())
		c.Header("Access-Control-Allow-Credentials", options.GetCredentials())
		c.Header("Access-Control-Max-Age", options.GetMaxAge())
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}
