package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"voo.su/internal/config"
	"voo.su/internal/delivery/http/handler"
	redisRepo "voo.su/internal/infrastructure/redis/repository"
	"voo.su/pkg/middleware"
	"voo.su/pkg/response"
)

func NewRouter(conf *config.Config, handler *handler.Handler, session *redisRepo.JwtTokenCacheRepository) *gin.Engine {
	router := gin.New()
	src, err := os.OpenFile(conf.App.LogPath("http_access.log"), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	router.Use(middleware.Cors(conf.App.Cors))
	router.Use(middleware.AccessLog(src))

	router.Use(gin.RecoveryWithWriter(gin.DefaultWriter, func(c *gin.Context, err any) {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Response{
			Code:    http.StatusInternalServerError,
			Message: "Ошибка системы, повторите попытку",
		})
	}))

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, response.Response{
			Code:    http.StatusOK,
			Message: "v1",
		})
	})

	NewV1(router, conf, handler, session)

	NewBot(router, handler)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, response.Response{
			Code:    http.StatusNotFound,
			Message: "Метод не найден",
		})
	})

	return router
}
