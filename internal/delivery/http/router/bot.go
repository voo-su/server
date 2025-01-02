// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package router

import (
	"github.com/gin-gonic/gin"
	"voo.su/internal/delivery/http/handler"
	"voo.su/pkg/core"
)

func NewBot(router *gin.Engine, handler *handler.Handler) {
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
