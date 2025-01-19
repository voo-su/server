package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"voo.su/internal/config"
	"voo.su/internal/constant"
	"voo.su/internal/delivery/ws/handler"
	"voo.su/internal/delivery/ws/middleware"
	redisRepo "voo.su/internal/infrastructure/redis/repository"
	"voo.su/pkg/ginutil"
	"voo.su/pkg/locale"
)

func NewRouter(conf *config.Config, locale locale.ILocale, handle *handler.Handler, session *redisRepo.JwtTokenCacheRepository) *gin.Engine {
	router := gin.New()
	src, err := os.OpenFile(conf.App.LogPath("ws_access.log"), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	router.Use(ginutil.AccessLog(src))

	router.Use(func(c *gin.Context) {
		acceptLang := c.GetHeader("Accept-Language")
		locale.SetFromHeaderAcceptLanguage(acceptLang)

		c.Next()
	})

	router.Use(gin.RecoveryWithWriter(gin.DefaultWriter, func(c *gin.Context, err any) {
		log.Println(err)

		c.AbortWithStatusJSON(http.StatusInternalServerError, ginutil.Response{
			Code:    http.StatusInternalServerError,
			Message: locale.Localize("network_error"),
		})
	}))

	authorize := middleware.Auth(locale, constant.GuardHttpAuth, conf.App.Jwt.Secret, session)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, ginutil.Response{
			Code:    http.StatusOK,
			Message: "success",
		})
	})

	router.GET("/ws", authorize, ginutil.HandlerFunc(handle.Chat.Conn))

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, ginutil.Response{
			Code:    http.StatusNotFound,
			Message: locale.Localize("method_not_found"),
		})
	})

	return router
}
