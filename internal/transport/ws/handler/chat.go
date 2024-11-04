package handler

import (
	"log"
	"voo.su/internal/repository/cache"
	"voo.su/internal/transport/ws/event"
	"voo.su/pkg/core"
	"voo.su/pkg/socket"
	"voo.su/pkg/socket/adapter"
)

type ChatChannel struct {
	Storage *cache.ClientStorage
	Event   *event.ChatEvent
}

func (c *ChatChannel) Conn(ctx *core.Context) error {
	conn, err := adapter.NewWsAdapter(ctx.Context.Writer, ctx.Context.Request)
	if err != nil {
		log.Printf("ошибка подключения к веб-сокету: %s", err.Error())
		return err
	}

	return c.NewClient(ctx.UserId(), conn)
}

func (c *ChatChannel) NewClient(uid int, conn socket.IConn) error {
	return socket.NewClient(conn, &socket.ClientOption{
		Uid:     uid,
		Channel: socket.Session.Chat,
		Storage: c.Storage,
		Buffer:  10,
	}, socket.NewEvent(
		socket.WithOpenEvent(c.Event.OnOpen),
		socket.WithMessageEvent(c.Event.OnMessage),
		socket.WithCloseEvent(c.Event.OnClose),
	))
}
