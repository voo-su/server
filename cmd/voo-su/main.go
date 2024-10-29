package main

import (
	"fmt"
	cliV2 "github.com/urfave/cli/v2"
	"voo.su/internal/cli"
	"voo.su/internal/config"
	"voo.su/internal/transport/http"
	"voo.su/internal/transport/ws"
	"voo.su/pkg/core"
	"voo.su/pkg/logger"
)

func NewHttpCommand() core.Command {
	return core.Command{
		Name: "http",
		Action: func(ctx *cliV2.Context, conf *config.Config) error {
			logger.InitLogger(fmt.Sprintf("%s/http.log", conf.App.Log), logger.LevelInfo, "http")

			return http.Run(ctx, NewHttpInjector(conf))
		},
	}
}

func NewWsCommand() core.Command {
	return core.Command{
		Name: "ws",
		Action: func(ctx *cliV2.Context, conf *config.Config) error {
			logger.InitLogger(fmt.Sprintf("%s/ws.log", conf.App.Log), logger.LevelInfo, "ws")

			return ws.Run(ctx, NewWsInjector(conf))
		},
	}
}

func NewCronCommand() core.Command {
	return core.Command{
		Name: "cli-cron",
		Action: func(ctx *cliV2.Context, conf *config.Config) error {
			logger.InitLogger(fmt.Sprintf("%s/cli_cron.log", conf.App.Log), logger.LevelInfo, "cron")

			return cli.Cron(ctx, NewCronInjector(conf))
		},
	}
}

func NewQueueCommand() core.Command {
	return core.Command{
		Name: "cli-queue",
		Action: func(ctx *cliV2.Context, conf *config.Config) error {
			logger.InitLogger(fmt.Sprintf("%s/cli_queue.log", conf.App.Log), logger.LevelInfo, "queue")

			return cli.Queue(ctx, NewQueueInjector(conf))
		},
	}
}

func NewMigrateCommand() core.Command {
	return core.Command{
		Name: "cli-migrate",
		Action: func(ctx *cliV2.Context, conf *config.Config) error {
			logger.InitLogger(fmt.Sprintf("%s/cli_migrate.log", conf.App.Log), logger.LevelInfo, "migrate")

			return cli.Migrate(ctx, NewMigrateInjector(conf))
		},
	}
}

func main() {
	app := core.NewApp()
	app.Register(NewHttpCommand())
	app.Register(NewWsCommand())
	app.Register(NewCronCommand())
	app.Register(NewQueueCommand())
	app.Register(NewMigrateCommand())
	app.Run()
}
