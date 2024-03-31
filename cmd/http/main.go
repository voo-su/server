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
	"voo.su/pkg/logger"
)

func main() {
	cmd := cli.NewApp()
	cmd.Name = "Voo.su"
	cmd.Usage = "HTTP-сервер"
	cmd.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "config",
			Aliases:     []string{"c"},
			Value:       "./init/voo-su.yaml",
			Usage:       "Путь к файлу конфигурации",
			DefaultText: "./init/voo-su.yaml",
		},
	}
	cmd.Action = func(tx *cli.Context) error {
		conf := config.New(tx.String("config"))
		logger.InitLogger(fmt.Sprintf("%s/http.log", conf.LogPath()), logger.LevelInfo, "http")
		if !conf.Debug() {
			gin.SetMode(gin.ReleaseMode)
		}
		app := Initialize(conf)
		eg, groupCtx := errgroup.WithContext(context.Background())
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
		log.Printf("HTTP Порт прослушивания: %d", conf.App.Http)
		log.Printf("HTTP PID сервера: %d", os.Getpid())
		return run(c, eg, groupCtx, app)
	}
	_ = cmd.Run(os.Args)
}

func run(c chan os.Signal, eg *errgroup.Group, ctx context.Context, app *AppProvider) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.Config.App.Http),
		Handler: app.Engine,
	}
	eg.Go(func() error {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})
	eg.Go(func() error {
		defer func() {
			log.Println("Выключение сервера...")
			timeCtx, timeCancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer timeCancel()
			if err := server.Shutdown(timeCtx); err != nil {
				log.Fatalf("Ошибка остановки HTTP-сервера: %s", err)
			}
		}()
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-c:
			return nil
		}
	})
	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		log.Fatalf("Принудительное завершение HTTP-сервера: %s", err)
	}
	log.Println("Выход из сервера")
	return nil
}
