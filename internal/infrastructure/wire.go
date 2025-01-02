// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

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
