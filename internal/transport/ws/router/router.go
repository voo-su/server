package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"voo.su/internal/config"
	"voo.su/internal/repository/cache"
	"voo.su/internal/transport/ws/handler"
	"voo.su/pkg/core"
	"voo.su/pkg/middleware"
	"voo.su/pkg/response"
)

func NewRouter(conf *config.Config, handle *handler.Handler, session *cache.JwtTokenStorage) *gin.Engine {
	router := gin.New()
	router.Use(gin.RecoveryWithWriter(gin.DefaultWriter, func(c *gin.Context, err any) {
		log.Println(err)

		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Response{
			Code:    http.StatusInternalServerError,
			Message: "Произошла ошибка системы. Пожалуйста, повторите попытку!",
		})
	}))

	authorize := middleware.Auth(conf.App.Jwt.Secret, "api", session)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, response.Response{
			Code:    http.StatusOK,
			Message: "success",
		})
	})

	router.GET("/ws", authorize, core.HandlerFunc(handle.Chat.Conn))

	//router.GET("/ws/detail", func(ctx *gin.Context) {
	//	ctx.JSON(http.StatusOK, map[string]any{
	//		"chat": socket.Session.Chat.Count(),
	//	})
	//})

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, response.Response{
			Code:    http.StatusNotFound,
			Message: "Метод не найден",
		})
	})

	return router
}
