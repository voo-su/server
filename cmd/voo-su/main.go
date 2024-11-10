package main

import (
	"fmt"
	cliV2 "github.com/urfave/cli/v2"
	"voo.su/internal/cli"
	"voo.su/internal/config"
	"voo.su/internal/provider"
	"voo.su/internal/transport/grpc"
	"voo.su/internal/transport/http"
	"voo.su/internal/transport/ws"
	"voo.su/pkg/logger"
)

func NewHttpCommand() provider.Command {
	return provider.Command{
		Name: "http",
		Action: func(ctx *cliV2.Context, conf *config.Config) error {
			logger.InitLogger(fmt.Sprintf("%s/http.log", conf.App.Log), logger.LevelInfo, "http")

			return http.Run(ctx, NewHttpInjector(conf))
		},
	}
}

func NewWsCommand() provider.Command {
	return provider.Command{
		Name: "ws",
		Action: func(ctx *cliV2.Context, conf *config.Config) error {
			logger.InitLogger(fmt.Sprintf("%s/ws.log", conf.App.Log), logger.LevelInfo, "ws")

			return ws.Run(ctx, NewWsInjector(conf))
		},
	}
}

func NewGrpcCommand() provider.Command {
	return provider.Command{
		Name: "grpc",
		Action: func(ctx *cliV2.Context, conf *config.Config) error {
			logger.InitLogger(fmt.Sprintf("%s/grpc.log", conf.App.Log), logger.LevelInfo, "grpc")

			return grpc.Run(ctx, NewGrpcInjector(conf))
		},
	}
}

func NewCronCommand() provider.Command {
	return provider.Command{
		Name: "cli-cron",
		Action: func(ctx *cliV2.Context, conf *config.Config) error {
			logger.InitLogger(fmt.Sprintf("%s/cli_cron.log", conf.App.Log), logger.LevelInfo, "cli-cron")

			return cli.Cron(ctx, NewCronInjector(conf))
		},
	}
}

func NewQueueCommand() provider.Command {
	return provider.Command{
		Name: "cli-queue",
		Action: func(ctx *cliV2.Context, conf *config.Config) error {
			logger.InitLogger(fmt.Sprintf("%s/cli_queue.log", conf.App.Log), logger.LevelInfo, "cli-queue")

			return cli.Queue(ctx, NewQueueInjector(conf))
		},
	}
}

func NewMigrateCommand() provider.Command {
	return provider.Command{
		Name: "cli-migrate",
		Action: func(ctx *cliV2.Context, conf *config.Config) error {
			logger.InitLogger(fmt.Sprintf("%s/cli_migrate.log", conf.App.Log), logger.LevelInfo, "cli-migrate")

			return cli.Migrate(ctx, NewMigrateInjector(conf))
		},
	}
}

func main() {
	app := provider.NewApp()
	app.Register(NewHttpCommand())
	app.Register(NewWsCommand())
	app.Register(NewCronCommand())
	app.Register(NewQueueCommand())
	app.Register(NewMigrateCommand())
	app.Register(NewGrpcCommand())
	app.Run()
}
