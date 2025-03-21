package clickhouse

import (
	"github.com/google/wire"
	"voo.su/internal/infrastructure/clickhouse/repository"
)

var ProviderSet = wire.NewSet(
	repository.NewAuthCodeRepository,
	repository.NewAccessLogRepository,
	repository.NewLoggerRepository,
	repository.NewAccessGrpcLogRepository,
	repository.NewAccessGrpcStreamLogRepository,
)
