package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"voo.su/internal/config"
	"voo.su/internal/repository/cache"
	"voo.su/internal/transport/ws/handler"
	"voo.su/pkg/core"
	"voo.su/pkg/core/middleware"
)

func NewRouter(conf *config.Config, handle *handler.Handler, session *cache.JwtTokenStorage) *gin.Engine {
	router := gin.New()

	router.Use(gin.RecoveryWithWriter(gin.DefaultWriter, func(c *gin.Context, err any) {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{"code": 500, "msg": "Произошла ошибка системы. Пожалуйста, повторите попытку!"})
	}))

	authorize := middleware.Auth(conf.Jwt.Secret, "api", session)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]any{"ok": "success"})
	})

	router.GET("/ws", authorize, core.HandlerFunc(handle.Chat.Conn))

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, map[string]any{"msg": "Метод не найден"})
	})

	return router
}
