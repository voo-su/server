//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"voo.su/internal/config"
	"voo.su/internal/logic"
	"voo.su/internal/provider"
	"voo.su/internal/repository/cache"
	"voo.su/internal/transport/http/handler"
	"voo.su/internal/transport/http/router"
)

var ProviderSet = wire.NewSet(
	wire.Struct(new(AppProvider), "*"),
	router.NewRouter,
	provider.NewPostgresqlClient,
	provider.NewRedisClient,
	provider.NewHttpClient,
	provider.NewEmailClient,
	provider.NewFilesystem,
	provider.NewRequestClient,

	wire.Struct(new(handler.Handler), "*"),
)

func Initialize(conf *config.Config) *AppProvider {
	panic(
		wire.Build(
			ProviderSet,
			cache.ProviderSet,
			logic.ProviderSet,
			handler.ProviderSet,
		),
	)
}
