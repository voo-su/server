package http

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
	"golang.org/x/sync/errgroup"
)

func Run(ctx *cli.Context, app *AppProvider) error {
	if app.Conf.App.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	eg, groupCtx := errgroup.WithContext(ctx.Context)

	c := make(chan os.Signal, 1)

	signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	log.Printf("HTTP Порт прослушивания: %d", app.Conf.App.Http)
	log.Printf("HTTP PID сервера: %d", os.Getpid())

	return run(c, eg, groupCtx, app)
}

func run(c chan os.Signal, eg *errgroup.Group, ctx context.Context, app *AppProvider) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.Conf.App.Http),
		Handler: app.Engine,
	}

	eg.Go(func() error {
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
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
