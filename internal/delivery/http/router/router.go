package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"voo.su/internal/config"
	"voo.su/internal/delivery/http/handler"
	"voo.su/internal/delivery/http/middleware"
	clickhouseRepo "voo.su/internal/infrastructure/clickhouse/repository"
	redisRepo "voo.su/internal/infrastructure/redis/repository"
	"voo.su/pkg/ginutil"
	"voo.su/pkg/locale"
)

func NewRouter(
	conf *config.Config,
	locale locale.ILocale,
	handler *handler.Handler,
	session *redisRepo.JwtTokenCacheRepository,
	accessLogRepo *clickhouseRepo.AccessLogRepository,
) *gin.Engine {
	router := gin.New()

	router.Use(ginutil.CorsMiddleware(
		conf.App.Cors.Origin,
		"Content-Type,User-Agent,Authorization",
		"OPTIONS,GET,POST,PUT,DELETE",
		conf.App.Cors.Credentials,
		conf.App.Cors.MaxAge,
	))
	router.Use(middleware.AccessLogMiddleware(accessLogRepo))

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

	NewManager(router, conf, handler)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, ginutil.Response{
			Code:    http.StatusNotFound,
			Message: locale.Localize("method_not_found"),
		})
	})

	return router
}
