package cli

import (
	"github.com/google/wire"
	"voo.su/internal/transport/cli/handle/cron"
)

var CronProviderSet = wire.NewSet(
	wire.Struct(new(CronProvider), "*"),
	wire.Struct(new(Crontab), "*"),

	cron.NewClearTmpFile,
	cron.NewClearWsCache,
	cron.NewClearExpireServer,
)
