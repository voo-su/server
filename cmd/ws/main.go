package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"voo.su/internal/config"
	"voo.su/pkg/core/socket"
	"voo.su/pkg/email"
	"voo.su/pkg/logger"
)

var ErrServerClosed = errors.New("закрытие сервера")

func main() {
	cmd := cli.NewApp()
	cmd.Name = "Voo.su"
	cmd.Usage = "WebSocket-сервер"
	cmd.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "config",
			Aliases:     []string{"c"},
			Value:       "./init/voo-su.yaml",
			Usage:       "Путь к файлу конфигурации",
			DefaultText: "./init/voo-su.yaml",
		},
	}
	cmd.Action = newApp
	_ = cmd.Run(os.Args)
}

func newApp(tx *cli.Context) error {
	eg, groupCtx := errgroup.WithContext(tx.Context)
	conf := config.New(tx.String("config"))
	logger.InitLogger(fmt.Sprintf("%s/ws.log", conf.LogPath()), logger.LevelInfo, "ws")
	if !conf.Debug() {
		gin.SetMode(gin.ReleaseMode)
	}
	app := Initialize(conf)
	socket.Initialize(groupCtx, eg, func(name string) {
		emailClient := app.Providers.EmailClient
		if conf.App.Env == "prod" {
			_ = emailClient.SendMail(&email.Option{
				To:      conf.Email.Report,
				Subject: fmt.Sprintf("%s Проблема с демоном", conf.App.Env),
				Body:    fmt.Sprintf("Проблема с демоном %s", name),
			})
		}
	})
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	time.AfterFunc(3*time.Second, func() {
		app.Coroutine.Start(eg, groupCtx)
	})
	log.Printf("ID сервера: %s", conf.ServerId())
	log.Printf("PID сервера: %d", os.Getpid())
	log.Printf("Порт прослушивания Websocket: %d", conf.App.Websocket)
	log.Printf("Порт прослушивания TCP: %d", conf.App.Tcp)
	go NewTcpServer(app)
	return start(c, eg, groupCtx, app)
}

func start(c chan os.Signal, eg *errgroup.Group, ctx context.Context, app *AppProvider) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.Config.App.Websocket),
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
