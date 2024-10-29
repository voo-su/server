package cli

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
	"os"
	"os/signal"
	"syscall"
	"voo.su/internal/cli/handler/queue"
	"voo.su/internal/config"
)

type QueueJobs struct {
	queue.EmailHandle
	queue.LoginHandle
}

type QueueProvider struct {
	Conf *config.Config
	DB   *gorm.DB
	Jobs *QueueJobs
}

func Queue(ctx *cli.Context, app *QueueProvider) error {
	if err := app.Jobs.EmailHandle.Handle(ctx.Context); err != nil {
		fmt.Println("EmailHandle>>", err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT)
	<-ch

	return nil
}
