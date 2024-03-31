//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"github.com/google/wire"
	"voo.su/internal/config"
	"voo.su/internal/provider"
	"voo.su/internal/repository/cache"
	"voo.su/internal/repository/repo"
	"voo.su/internal/transport/cli/command"
	"voo.su/internal/transport/cli/command/cron"
	cron2 "voo.su/internal/transport/cli/handle/cron"
	"voo.su/pkg/filesystem"
)

func Initialize(ctx context.Context, conf *config.Config) *AppProvider {
	panic(wire.Build(wire.NewSet(
		wire.Struct(new(AppProvider), "*"),

		provider.NewPostgresqlClient,
		provider.NewRedisClient,
		provider.NewRequestClient,

		filesystem.NewFilesystem,

		wire.Struct(new(command.Commands), "*"),

		cron.NewCrontabCommand,
		cron2.NewClearTmpFile,
		cron2.NewClearWsCache,
		cron2.NewClearExpireServer,

		cache.NewSidStorage,
		cache.NewSequence,

		repo.NewSource,
		repo.NewSequence,
	)))
}
