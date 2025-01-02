// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package clickhouse

import (
	"github.com/google/wire"
	"voo.su/internal/infrastructure/clickhouse/repository"
)

var ProviderSet = wire.NewSet(
	repository.NewAuthCodeRepository,
)
