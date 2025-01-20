package router

import (
	"github.com/gin-gonic/gin"
	"voo.su/internal/delivery/http/handler"
	"voo.su/pkg/ginutil"
)

func NewManager(router *gin.Engine, handler *handler.Handler) {
	bot := router.Group("/manager")
	{
		bot.GET("/dashboard", ginutil.HandlerFunc(handler.Manager.Dashboard.Dashboard))
	}
}
