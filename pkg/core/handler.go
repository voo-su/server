package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"voo.su/pkg/response"
)

func HandlerFunc(fn func(ctx *Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := fn(New(c)); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, &response.Response{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
				//Meta:    initMeta(),
			})

			return
		}
	}
}
