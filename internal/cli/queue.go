package cli

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
	"log"
	"os"
	"os/signal"
	"syscall"
	"voo.su/internal/cli/handler/queue"
	"voo.su/internal/config"
	clickhouseRepo "voo.su/internal/infrastructure/clickhouse/repository"
	"voo.su/internal/provider"
)

type QueueJobs struct {
	queue.EmailHandle
	queue.PushHandle
}

type QueueProvider struct {
	Conf             *config.Config
	DB               *gorm.DB
	Jobs             *QueueJobs
	LoggerRepository *clickhouseRepo.LoggerRepository
}

func Queue(ctx *cli.Context, app *QueueProvider) error {
	log.SetOutput(provider.NewLoggerWriter(app.Conf, os.Stdout, app.LoggerRepository))

	if err := app.Jobs.EmailHandle.Handle(ctx.Context); err != nil {
		fmt.Println("EmailHandle>>", err)
	}

	if err := app.Jobs.PushHandle.Handle(ctx.Context); err != nil {
		fmt.Println("PushHandle>>", err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT)
	<-ch

	return nil
}
