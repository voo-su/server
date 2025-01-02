// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package socket

import (
	"log"
	"voo.su/pkg/utils"
)

type IEvent interface {
	Open(client IClient)

	Message(client IClient, data []byte)

	Close(client IClient, code int, text string)

	Destroy(client IClient)
}

type (
	OpenEvent    func(client IClient)
	MessageEvent func(client IClient, data []byte)
	CloseEvent   func(client IClient, code int, text string)
	DestroyEvent func(client IClient)
	EventOption  func(event *Event)
)

type Event struct {
	open    OpenEvent
	message MessageEvent
	close   CloseEvent
	destroy DestroyEvent
}

func NewEvent(opts ...EventOption) IEvent {
	o := &Event{}
	for _, opt := range opts {
		opt(o)
	}

	return o
}

func (e *Event) Open(client IClient) {
	if e.open == nil {
		return
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println("Исключение обратного вызова события 'open': ", client.Uid(), client.Cid(), client.Channel().Name(), utils.PanicTrace(err))
		}
	}()

	e.open(client)
}

func (e *Event) Message(client IClient, data []byte) {
	if e.message == nil {
		return
	}
	defer func() {
		if err := recover(); err != nil {
			log.Println("Исключение обратного вызова события 'message': ", client.Uid(), client.Cid(), client.Channel().Name(), utils.PanicTrace(err))
		}
	}()

	e.message(client, data)
}

func (e *Event) Close(client IClient, code int, text string) {
	if e.close == nil {
		return
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println("Исключение обратного вызова события 'close': ", client.Uid(), client.Cid(), client.Channel().Name(), utils.PanicTrace(err))
		}
	}()

	e.close(client, code, text)
}

func (e *Event) Destroy(client IClient) {
	if e.destroy == nil {
		return
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println("Исключение обратного вызова события 'destroy': ", client.Uid(), client.Cid(), client.Channel().Name(), utils.PanicTrace(err))
		}
	}()

	e.destroy(client)
}

func WithOpenEvent(e OpenEvent) EventOption {
	return func(event *Event) {
		event.open = e
	}
}

func WithMessageEvent(e MessageEvent) EventOption {
	return func(event *Event) {
		event.message = e
	}
}

func WithCloseEvent(e CloseEvent) EventOption {
	return func(event *Event) {
		event.close = e
	}
}

func WithDestroyEvent(e DestroyEvent) EventOption {
	return func(event *Event) {
		event.destroy = e
	}
}
