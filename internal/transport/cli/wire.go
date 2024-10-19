package cli

import (
	"github.com/google/wire"
	"voo.su/internal/transport/cli/handle/cron"
	"voo.su/internal/transport/cli/handle/queue"
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
	wire.Struct(new(queue.LoginHandle), "*"),
)

var MigrateProviderSet = wire.NewSet(
	wire.Struct(new(MigrateProvider), "*"),
)
