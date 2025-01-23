package router

import (
	"github.com/gin-gonic/gin"
	"voo.su/internal/config"
	"voo.su/internal/delivery/http/handler"
	"voo.su/pkg/ginutil"
)

func NewManager(router *gin.Engine, conf *config.Config, handler *handler.Handler) {
	manager := router.Group("/manager")
	{
		manager.Use(ginutil.IPWhitelistMiddleware(conf.Manager.Ips))

		manager.GET("/dashboard", ginutil.HandlerFunc(handler.Manager.Dashboard.Dashboard))
	}
}
