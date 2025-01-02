// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package handler

import (
	"voo.su/internal/delivery/http/handler/bot"
	"voo.su/internal/delivery/http/handler/v1"
)

type Handler struct {
	V1  *v1.Handler
	Bot *bot.Handler
}
