package infrastructure

import (
	"github.com/google/wire"
	"voo.su/internal/infrastructure/clickhouse"
	"voo.su/internal/infrastructure/postgres"
	"voo.su/internal/infrastructure/redis"
)

var ProviderSet = wire.NewSet(
	NewSource,
	postgres.ProviderSet,
	redis.ProviderSet,
	clickhouse.ProviderSet,
)
