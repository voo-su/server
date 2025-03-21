//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"voo.su/internal/cli"
	"voo.su/internal/config"
	"voo.su/internal/delivery/grpc"
	"voo.su/internal/delivery/http"
	"voo.su/internal/delivery/ws"
	"voo.su/internal/provider"
)

func NewHttpInjector(conf *config.Config) *http.AppProvider {
	panic(wire.Build(provider.ProviderSet, http.ProviderSet))
}

func NewWsInjector(conf *config.Config) *ws.AppProvider {
	panic(wire.Build(provider.ProviderSet, ws.ProviderSet))
}

func NewGrpcInjector(conf *config.Config) *grpc.AppProvider {
	panic(wire.Build(provider.ProviderSet, grpc.ProviderSet))
}

func NewCronInjector(conf *config.Config) *cli.CronProvider {
	panic(wire.Build(provider.ProviderSet, cli.CronProviderSet))
}

func NewQueueInjector(conf *config.Config) *cli.QueueProvider {
	panic(wire.Build(provider.ProviderSet, cli.QueueProviderSet))
}

func NewMigrateInjector(conf *config.Config) *cli.MigrateProvider {
	panic(wire.Build(provider.ProviderSet, cli.MigrateProviderSet))
}

func NewGenerateInjector(conf *config.Config) *cli.GenerateProvider {
	panic(wire.Build(provider.ProviderSet, cli.GenerateProviderSet))
}
