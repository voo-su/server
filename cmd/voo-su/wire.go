//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"voo.su/internal/cli"
	"voo.su/internal/config"
	"voo.su/internal/domain/logic"
	"voo.su/internal/provider"
	"voo.su/internal/repository"
	"voo.su/internal/repository/cache"
	"voo.su/internal/transport/grpc"
	"voo.su/internal/transport/http"
	"voo.su/internal/transport/ws"
	"voo.su/internal/usecase"
)

var providerSet = wire.NewSet(
	provider.NewPostgresqlClient,
	provider.NewClickHouseClient,
	provider.NewRedisClient,
	provider.NewHttpClient,
	provider.NewEmailClient,
	provider.NewFilesystem,
	provider.NewRequestClient,

	wire.Struct(new(provider.Providers), "*"),

	cache.ProviderSet,
	logic.ProviderSet,
	usecase.ProviderSet,
	repository.ProviderSet,
)

func NewHttpInjector(conf *config.Config) *http.AppProvider {
	panic(
		wire.Build(
			providerSet,
			http.ProviderSet,
		),
	)
}

func NewWsInjector(conf *config.Config) *ws.AppProvider {
	panic(
		wire.Build(
			providerSet,
			ws.ProviderSet,
		),
	)
}

func NewCronInjector(conf *config.Config) *cli.CronProvider {
	panic(
		wire.Build(
			providerSet,
			cli.CronProviderSet,
		),
	)
}

func NewQueueInjector(conf *config.Config) *cli.QueueProvider {
	panic(
		wire.Build(
			providerSet,
			cli.QueueProviderSet,
		),
	)
}

func NewMigrateInjector(conf *config.Config) *cli.MigrateProvider {
	panic(
		wire.Build(
			providerSet,
			cli.MigrateProviderSet,
		),
	)
}

func NewGrpcInjector(conf *config.Config) *grpc.AppProvider {
	panic(wire.Build(
		providerSet,
		grpc.ProviderSet,
	))
}
