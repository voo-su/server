package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"voo.su/internal/config"
	"voo.su/internal/delivery/http/handler"
	redisRepo "voo.su/internal/infrastructure/redis/repository"
	"voo.su/pkg/ginutil"
	"voo.su/pkg/locale"
)

func NewRouter(conf *config.Config, locale locale.ILocale, handler *handler.Handler, session *redisRepo.JwtTokenCacheRepository) *gin.Engine {
	router := gin.New()
	src, err := os.OpenFile(conf.App.LogPath("http_access.log"), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	router.Use(ginutil.Cors(conf.App.Cors))
	router.Use(ginutil.AccessLog(src))

	router.Use(func(c *gin.Context) {
		acceptLang := c.GetHeader("Accept-Language")
		locale.SetFromHeaderAcceptLanguage(acceptLang)

		c.Next()
	})

	router.Use(gin.RecoveryWithWriter(gin.DefaultWriter, func(c *gin.Context, err any) {
		c.AbortWithStatusJSON(http.StatusInternalServerError, ginutil.Response{
			Code:    http.StatusInternalServerError,
			Message: locale.Localize("network_error"),
		})
	}))

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, ginutil.Response{
			Code:    http.StatusOK,
			Message: "v1",
		})
	})

	NewV1(conf, locale, router, handler, session)

	NewBot(router, handler)

	NewManager(router, handler)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, ginutil.Response{
			Code:    http.StatusNotFound,
			Message: locale.Localize("method_not_found"),
		})
	})

	return router
}
