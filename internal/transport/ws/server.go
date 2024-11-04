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
	"voo.su/pkg/socket"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
	"golang.org/x/sync/errgroup"
	"voo.su/internal/config"
	"voo.su/internal/provider"
	"voo.su/internal/transport/ws/handler"
	"voo.su/internal/transport/ws/process"
	"voo.su/pkg/email"
)

var ErrServerClosed = errors.New("закрытие сервера")

type AppProvider struct {
	Conf      *config.Config
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
				Subject: fmt.Sprintf("%s Проблема с демоном", app.Conf.App.Env),
				Body:    fmt.Sprintf("Проблема с демоном %s", name),
			})
		}
	})

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	time.AfterFunc(3*time.Second, func() {
		app.Coroutine.Start(eg, groupCtx)
	})

	log.Printf("ID сервера: %s", app.Conf.ServerId())
	log.Printf("PID сервера: %d", os.Getpid())
	log.Printf("Порт прослушивания Websocket: %d", app.Conf.Server.Websocket)
	log.Printf("Порт прослушивания TCP: %d", app.Conf.Server.Tcp)

	go NewTcpServer(app)

	return start(c, eg, groupCtx, app)
}

func start(c chan os.Signal, eg *errgroup.Group, ctx context.Context, app *AppProvider) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.Conf.Server.Websocket),
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
			log.Println("Выключение сервера...")
			timeCtx, timeCancel := context.WithTimeout(context.TODO(), 3*time.Second)
			defer timeCancel()
			if err := server.Shutdown(timeCtx); err != nil {
				log.Printf("Ошибка остановки Websocket-сервера: %s \n", err)
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
		log.Fatalf("Принудительное завершение сервера: %s", err)
	}

	log.Println("Выход из сервера")

	return nil
}

func NewTcpServer(app *AppProvider) {
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", app.Conf.Server.Tcp))
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = listener.Close()
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Ошибка при принятии соединения:", err)
			continue
		}

		go app.Handler.Dispatch(conn)
	}
}
