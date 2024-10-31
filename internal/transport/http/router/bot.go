package router

import (
	"github.com/gin-gonic/gin"
	"voo.su/internal/config"
	"voo.su/internal/repository/cache"
	"voo.su/internal/transport/http/handler"
	"voo.su/pkg/core"
)

func NewBot(router *gin.Engine, conf *config.Config, handler *handler.Handler, session *cache.JwtTokenStorage) {
	bot := router.Group("/bot/:token")
	{
		bot.POST("/send", core.HandlerFunc(handler.Bot.Message.Send))
		bot.GET("/group-chats", core.HandlerFunc(handler.Bot.Message.GroupChats))
	}
}
