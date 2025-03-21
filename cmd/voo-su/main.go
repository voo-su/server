package main

import (
	cliV2 "github.com/urfave/cli/v2"
	"voo.su/internal/cli"
	"voo.su/internal/config"
	"voo.su/internal/delivery/grpc"
	"voo.su/internal/delivery/http"
	"voo.su/internal/delivery/ws"
	"voo.su/internal/provider"
)

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

func NewHttpCommand() provider.Command {
	return provider.Command{
		Name: "http",
		Action: func(ctx *cliV2.Context, conf *config.Config) error {
			return http.Run(ctx, NewHttpInjector(conf))
		},
	}
}

func NewWsCommand() provider.Command {
	return provider.Command{
		Name: "ws",
		Action: func(ctx *cliV2.Context, conf *config.Config) error {
			return ws.Run(ctx, NewWsInjector(conf))
		},
	}
}

func NewGrpcCommand() provider.Command {
	return provider.Command{
		Name: "grpc",
		Action: func(ctx *cliV2.Context, conf *config.Config) error {
			return grpc.Run(ctx, NewGrpcInjector(conf))
		},
	}
}

func NewCronCommand() provider.Command {
	return provider.Command{
		Name: "cli-cron",
		Action: func(ctx *cliV2.Context, conf *config.Config) error {
			return cli.Cron(ctx, NewCronInjector(conf))
		},
	}
}

func NewQueueCommand() provider.Command {
	return provider.Command{
		Name: "cli-queue",
		Action: func(ctx *cliV2.Context, conf *config.Config) error {
			return cli.Queue(ctx, NewQueueInjector(conf))
		},
	}
}

func NewMigrateCommand() provider.Command {
	return provider.Command{
		Name: "cli-migrate",
		Action: func(ctx *cliV2.Context, conf *config.Config) error {
			return cli.Migrate(ctx, NewMigrateInjector(conf))
		},
	}
}

func NewGenerateCommand() provider.Command {
	return provider.Command{
		Name: "cli-generate",
		Action: func(ctx *cliV2.Context, conf *config.Config) error {
			return cli.Generate(ctx, NewGenerateInjector(conf))
		},
	}
}
