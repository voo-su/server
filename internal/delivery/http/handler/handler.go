package handler

import (
	"voo.su/internal/delivery/http/handler/bot"
	"voo.su/internal/delivery/http/handler/v1"
)

type Handler struct {
	V1  *v1.Handler
	Bot *bot.Handler
}
