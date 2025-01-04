package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"voo.su/internal/config"
	"voo.su/internal/constant"
	"voo.su/internal/delivery/ws/handler"
	redisRepo "voo.su/internal/infrastructure/redis/repository"
	"voo.su/pkg/core"
	"voo.su/pkg/locale"
	"voo.su/pkg/middleware"
	"voo.su/pkg/response"
)

func NewRouter(conf *config.Config, locale locale.ILocale, handle *handler.Handler, session *redisRepo.JwtTokenCacheRepository) *gin.Engine {
	router := gin.New()
	src, err := os.OpenFile(conf.App.LogPath("ws_access.log"), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	router.Use(middleware.AccessLog(src))

	router.Use(gin.RecoveryWithWriter(gin.DefaultWriter, func(c *gin.Context, err any) {
		log.Println(err)

		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Response{
			Code:    http.StatusInternalServerError,
			Message: locale.Localize("network_error"),
		})
	}))

	authorize := middleware.Auth(locale, constant.GuardHttpAuth, conf.App.Jwt.Secret, session)

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
			Message: locale.Localize("method_not_found"),
		})
	})

	return router
}
