package ws

import (
	"github.com/google/wire"
	"voo.su/internal/repository"
	"voo.su/internal/transport/ws/consume"
	"voo.su/internal/transport/ws/event"
	"voo.su/internal/transport/ws/handler"
	"voo.su/internal/transport/ws/process"
	"voo.su/internal/transport/ws/router"
	"voo.su/internal/usecase"
)

var ProviderSet = wire.NewSet(
	wire.Struct(new(handler.Handler), "*"),
	wire.Struct(new(AppProvider), "*"),
	wire.Struct(new(process.SubServers), "*"),

	router.NewRouter,
	handler.ProviderSet,
	event.ProviderSet,
	consume.ProviderSet,
	process.ProviderSet,
	usecase.ProviderSet,
	repository.ProviderSet,
)
