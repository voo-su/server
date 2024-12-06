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
		bot.GET("/group-chats", core.HandlerFunc(handler.Bot.Message.GroupChats))
		bot.POST("/send/message", core.HandlerFunc(handler.Bot.Message.Message))
		bot.POST("/send/photo", core.HandlerFunc(handler.Bot.Message.Photo))
		bot.POST("/send/video", core.HandlerFunc(handler.Bot.Message.Video))
		bot.POST("/send/audio", core.HandlerFunc(handler.Bot.Message.Audio))
		bot.POST("/send/document", core.HandlerFunc(handler.Bot.Message.Document))
	}
}
