// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package ws

import (
	"github.com/google/wire"
	"voo.su/internal/delivery/ws/consume"
	"voo.su/internal/delivery/ws/event"
	"voo.su/internal/delivery/ws/handler"
	"voo.su/internal/delivery/ws/process"
	"voo.su/internal/delivery/ws/router"
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
)
