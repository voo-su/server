package main

import (
	"github.com/gin-gonic/gin"
	"voo.su/internal/config"
)

type AppProvider struct {
	Config *config.Config
	Engine *gin.Engine
}
