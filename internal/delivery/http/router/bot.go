package router

import (
	"github.com/gin-gonic/gin"
	"voo.su/internal/delivery/http/handler"
	"voo.su/pkg/ginutil"
)

func NewBot(router *gin.Engine, handler *handler.Handler) {
	bot := router.Group("/bot/:token")
	{
		bot.GET("/group-chats", ginutil.HandlerFunc(handler.Bot.Message.GroupChats))
		bot.POST("/send/message", ginutil.HandlerFunc(handler.Bot.Message.Message))
		bot.POST("/send/photo", ginutil.HandlerFunc(handler.Bot.Message.Photo))
		bot.POST("/send/video", ginutil.HandlerFunc(handler.Bot.Message.Video))
		bot.POST("/send/audio", ginutil.HandlerFunc(handler.Bot.Message.Audio))
		bot.POST("/send/document", ginutil.HandlerFunc(handler.Bot.Message.Document))
	}
}
