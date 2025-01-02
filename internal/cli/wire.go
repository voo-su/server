// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package cli

import (
	"github.com/google/wire"
	"voo.su/internal/cli/handler/cron"
	"voo.su/internal/cli/handler/queue"
)

var CronProviderSet = wire.NewSet(
	wire.Struct(new(CronProvider), "*"),
	wire.Struct(new(Crontab), "*"),

	cron.NewClearTmpFile,
	cron.NewClearWsCache,
	cron.NewClearExpireServer,
)

var QueueProviderSet = wire.NewSet(
	wire.Struct(new(QueueProvider), "*"),
	wire.Struct(new(QueueJobs), "*"),
	wire.Struct(new(queue.EmailHandle), "*"),
	wire.Struct(new(queue.PushHandle), "*"),
)

var MigrateProviderSet = wire.NewSet(
	wire.Struct(new(MigrateProvider), "*"),
)

var GenerateProviderSet = wire.NewSet(
	wire.Struct(new(GenerateProvider), "*"),
)
