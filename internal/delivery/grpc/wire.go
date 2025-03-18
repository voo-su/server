package grpc

import (
	"github.com/google/wire"
	"voo.su/internal/delivery/grpc/handler"
	"voo.su/internal/delivery/grpc/middleware"
)

var ProviderSet = wire.NewSet(
	wire.Struct(new(AppProvider), "*"),
	handler.ProviderSet,
	middleware.ProviderSet,
)
