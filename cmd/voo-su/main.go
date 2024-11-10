package main

import (
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
			logger.InitLogger(conf.App.LogPath("http.log"), logger.LevelInfo, "http")

			return http.Run(ctx, NewHttpInjector(conf))
		},
	}
}

func NewWsCommand() provider.Command {
	return provider.Command{
		Name: "ws",
		Action: func(ctx *cliV2.Context, conf *config.Config) error {
			logger.InitLogger(conf.App.LogPath("ws.log"), logger.LevelInfo, "ws")

			return ws.Run(ctx, NewWsInjector(conf))
		},
	}
}

func NewGrpcCommand() provider.Command {
	return provider.Command{
		Name: "grpc",
		Action: func(ctx *cliV2.Context, conf *config.Config) error {
			logger.InitLogger(conf.App.LogPath("grpc.log"), logger.LevelInfo, "grpc")

			return grpc.Run(ctx, NewGrpcInjector(conf))
		},
	}
}

func NewCronCommand() provider.Command {
	return provider.Command{
		Name: "cli-cron",
		Action: func(ctx *cliV2.Context, conf *config.Config) error {
			logger.InitLogger(conf.App.LogPath("cli-cron.log"), logger.LevelInfo, "cli-cron")

			return cli.Cron(ctx, NewCronInjector(conf))
		},
	}
}

func NewQueueCommand() provider.Command {
	return provider.Command{
		Name: "cli-queue",
		Action: func(ctx *cliV2.Context, conf *config.Config) error {
			logger.InitLogger(conf.App.LogPath("cli-queue.log"), logger.LevelInfo, "cli-queue")

			return cli.Queue(ctx, NewQueueInjector(conf))
		},
	}
}

func NewMigrateCommand() provider.Command {
	return provider.Command{
		Name: "cli-migrate",
		Action: func(ctx *cliV2.Context, conf *config.Config) error {
			logger.InitLogger(conf.App.LogPath("cli-migrate.log"), logger.LevelInfo, "cli-migrate")

			return cli.Migrate(ctx, NewMigrateInjector(conf))
		},
	}
}

func NewGenerateCommand() provider.Command {
	return provider.Command{
		Name: "cli-generate",
		Action: func(ctx *cliV2.Context, conf *config.Config) error {
			logger.InitLogger(conf.App.LogPath("cli-generate.log"), logger.LevelInfo, "cli-generate")

			return cli.Generate(ctx, NewGenerateInjector(conf))
		},
	}
}

func main() {
	app := provider.NewApp()
	app.Register(NewHttpCommand())
	app.Register(NewWsCommand())
	app.Register(NewGrpcCommand())
	app.Register(NewCronCommand())
	app.Register(NewQueueCommand())
	app.Register(NewMigrateCommand())
	app.Register(NewGenerateCommand())
	app.Run()
}
