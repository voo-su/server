package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"voo.su/internal/config"
	"voo.su/internal/repository/cache"
	"voo.su/internal/transport/http/handler"
	"voo.su/pkg/core/middleware"
)

func NewRouter(conf *config.Config, handler *handler.Handler, session *cache.JwtTokenStorage) *gin.Engine {
	router := gin.New()
	src, err := os.OpenFile(fmt.Sprintf("%s/http_access.log", conf.App.Log), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	router.Use(middleware.Cors(conf.Cors))
	router.Use(middleware.AccessLog(src))

	router.Use(gin.RecoveryWithWriter(gin.DefaultWriter, func(c *gin.Context, err any) {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
			"code":    http.StatusInternalServerError,
			"message": "Ошибка системы, повторите попытку",
		})
	}))

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]any{
			"code":    http.StatusOK,
			"message": "v1",
		})
	})

	NewV1(router, conf, handler, session)

	NewBot(router, conf, handler, session)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, map[string]any{
			"code":    http.StatusNotFound,
			"message": "Метод не найден",
		})
	})

	return router
}
