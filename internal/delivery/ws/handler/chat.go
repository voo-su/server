// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package handler

import (
	"log"
	"voo.su/internal/delivery/ws/event"
	redisRepo "voo.su/internal/infrastructure/redis/repository"
	"voo.su/pkg/core"
	"voo.su/pkg/socket"
	"voo.su/pkg/socket/adapter"
)

type ChatChannel struct {
	ClientCacheRepo *redisRepo.ClientCacheRepository
	Event           *event.ChatEvent
}

func (c *ChatChannel) Conn(ctx *core.Context) error {
	conn, err := adapter.NewWsAdapter(ctx.Context.Writer, ctx.Context.Request)
	if err != nil {
		log.Printf("WS Conn error: %s", err.Error())
		return err
	}

	return c.NewClient(ctx.UserId(), conn)
}

func (c *ChatChannel) NewClient(uid int, conn socket.IConn) error {
	return socket.NewClient(conn, &socket.ClientOption{
		Uid:     uid,
		Channel: socket.Session.Chat,
		Storage: c.ClientCacheRepo,
		Buffer:  10,
	}, socket.NewEvent(
		socket.WithOpenEvent(c.Event.OnOpen),
		socket.WithMessageEvent(c.Event.OnMessage),
		socket.WithCloseEvent(c.Event.OnClose),
	))
}
