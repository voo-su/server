package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"voo.su/internal/config"
	"voo.su/internal/provider"
	"voo.su/internal/transport/ws/handler"
	"voo.su/internal/transport/ws/process"
)

type AppProvider struct {
	Config    *config.Config
	Engine    *gin.Engine
	Coroutine *process.Server
	Handler   *handler.Handler
	Providers *provider.Providers
}

func NewTcpServer(app *AppProvider) {
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", app.Config.App.Tcp))
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = listener.Close()
	}()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("не удалось принять соединение, ошибка:", err)
			continue
		}
		go app.Handler.Dispatch(conn)
	}
}
