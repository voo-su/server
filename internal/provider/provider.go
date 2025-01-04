package provider

import (
	"github.com/google/wire"
	"voo.su/internal/domain/logic"
	"voo.su/internal/infrastructure"
	"voo.su/internal/usecase"
	"voo.su/pkg/email"
)

type Providers struct {
	EmailClient *email.Email
}

var ProviderSet = wire.NewSet(
	NewPostgresqlClient,
	NewClickHouseClient,
	NewRedisClient,
	NewHttpClient,
	NewEmailClient,
	NewMinioClient,
	NewRequestClient,
	NewNatsClient,
	NewLocale,

	wire.Struct(new(Providers), "*"),

	logic.ProviderSet,
	usecase.ProviderSet,
	infrastructure.ProviderSet,
)
