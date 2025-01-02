// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package ws

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"voo.su/pkg/locale"
	"voo.su/pkg/socket"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
	"golang.org/x/sync/errgroup"
	"voo.su/internal/config"
	"voo.su/internal/delivery/ws/handler"
	"voo.su/internal/delivery/ws/process"
	"voo.su/internal/provider"
	"voo.su/pkg/email"
)

var ErrServerClosed = errors.New("server shutdown")

type AppProvider struct {
	Conf      *config.Config
	Locale    locale.ILocale
	Engine    *gin.Engine
	Coroutine *process.Server
	Handler   *handler.Handler
	Providers *provider.Providers
}

func Run(ctx *cli.Context, app *AppProvider) error {
	eg, groupCtx := errgroup.WithContext(ctx.Context)

	if app.Conf.App.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	socket.Initialize(groupCtx, eg, func(name string) {
		emailClient := app.Providers.EmailClient
		if app.Conf.App.Env == "prod" {
			_ = emailClient.SendMail(&email.Option{
				To:      app.Conf.Email.Report,
				Subject: fmt.Sprintf(app.Locale.Localize("service_start"), app.Conf.App.Env),
				Body:    fmt.Sprintf(app.Locale.Localize("service_start"), name),
			})
		}
	})

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	time.AfterFunc(3*time.Second, func() {
		app.Coroutine.Start(eg, groupCtx)
	})

	log.Printf("ID server: %s", app.Conf.ServerId())
	log.Printf("PID server: %d", os.Getpid())
	log.Printf("Websocket: %s", app.Conf.Server.Websocket.GetWebsocket())
	log.Printf("TCP: %s", app.Conf.Server.Tcp.GetTcp())

	go NewTcpServer(app)

	return start(c, eg, groupCtx, app)
}

func start(c chan os.Signal, eg *errgroup.Group, ctx context.Context, app *AppProvider) error {
	server := &http.Server{
		Addr:    app.Conf.Server.Websocket.GetWebsocket(),
		Handler: app.Engine,
	}
	eg.Go(func() error {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}

		return nil
	})
	eg.Go(func() (err error) {
		defer func() {
			log.Println("Shutting down server...")
			timeCtx, timeCancel := context.WithTimeout(context.TODO(), 3*time.Second)
			defer timeCancel()
			if err := server.Shutdown(timeCtx); err != nil {
				log.Printf("Error stopping Websocket server: %s \n", err)
			}
			err = ErrServerClosed
		}()

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-c:
			return nil
		}
	})

	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) && !errors.Is(err, ErrServerClosed) {
		log.Fatalf("Forced server shutdown: %s", err)
	}

	return nil
}

func NewTcpServer(app *AppProvider) {
	listener, err := net.Listen("tcp", app.Conf.Server.Tcp.GetTcp())
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = listener.Close()
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %s", err)
			continue
		}

		go app.Handler.Dispatch(conn)
	}
}
