// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package http

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"voo.su/internal/config"
	"voo.su/internal/delivery/http/handler"
	"voo.su/internal/delivery/http/handler/bot"
	v1 "voo.su/internal/delivery/http/handler/v1"
	"voo.su/internal/delivery/http/router"
)

type AppProvider struct {
	Conf   *config.Config
	Engine *gin.Engine
}

var ProviderSet = wire.NewSet(
	wire.Struct(new(AppProvider), "*"),
	router.NewRouter,

	wire.Struct(new(handler.Handler), "*"),

	// v1
	wire.Struct(new(v1.Handler), "*"),
	wire.Struct(new(v1.Auth), "*"),
	wire.Struct(new(v1.Account), "*"),
	wire.Struct(new(v1.Upload), "*"),
	wire.Struct(new(v1.Contact), "*"),
	wire.Struct(new(v1.ContactRequest), "*"),
	wire.Struct(new(v1.Chat), "*"),
	wire.Struct(new(v1.Message), "*"),
	wire.Struct(new(v1.GroupChat), "*"),
	wire.Struct(new(v1.GroupChatRequest), "*"),
	wire.Struct(new(v1.Sticker), "*"),
	wire.Struct(new(v1.ContactFolder), "*"),
	wire.Struct(new(v1.GroupChatAds), "*"),
	wire.Struct(new(v1.Search), "*"),
	wire.Struct(new(v1.Bot), "*"),
	wire.Struct(new(v1.Project), "*"),
	wire.Struct(new(v1.ProjectTask), "*"),
	wire.Struct(new(v1.ProjectTaskComment), "*"),

	// Bot
	wire.Struct(new(bot.Handler), "*"),
	wire.Struct(new(bot.Message), "*"),
)
