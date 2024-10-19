//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"voo.su/internal/config"
	"voo.su/internal/logic"
	"voo.su/internal/provider"
	"voo.su/internal/repository/cache"
	"voo.su/internal/transport/cli"
	"voo.su/internal/transport/http"
	"voo.su/internal/transport/ws"
)

var providerSet = wire.NewSet(
	provider.NewPostgresqlClient,
	provider.NewRedisClient,
	provider.NewHttpClient,
	provider.NewEmailClient,
	provider.NewFilesystem,
	provider.NewRequestClient,

	wire.Struct(new(provider.Providers), "*"),

	cache.ProviderSet,
	logic.ProviderSet,
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
