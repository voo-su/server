// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package process

import (
	"context"
	"golang.org/x/sync/errgroup"
	"reflect"
	"sync"
)

var once sync.Once

type IServer interface {
	Setup(ctx context.Context) error
}

type SubServers struct {
	HealthSubscribe  *HealthSubscribe
	MessageSubscribe *MessageSubscribe
}

type Server struct {
	items []IServer
}

func NewServer(servers *SubServers) *Server {
	s := &Server{}
	s.binds(servers)
	return s
}

func (c *Server) binds(servers *SubServers) {
	elem := reflect.ValueOf(servers).Elem()
	for i := 0; i < elem.NumField(); i++ {
		if v, ok := elem.Field(i).Interface().(IServer); ok {
			c.items = append(c.items, v)
		}
	}
}

func (c *Server) Start(eg *errgroup.Group, ctx context.Context) {
	once.Do(func() {
		for _, process := range c.items {
			func(serv IServer) {
				eg.Go(func() error {
					return serv.Setup(ctx)
				})
			}(process)
		}
	})
}
