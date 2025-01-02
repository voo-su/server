// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package nats

import (
	"fmt"
	"github.com/nats-io/nats.go"
)

type INatsClient interface {
	Publish(subject string, data []byte) error

	Subscribe(subject string, handler func(msg *nats.Msg)) (*nats.Subscription, error)

	Close()
}

var _ INatsClient = (*NatsClient)(nil)

type NatsClient struct {
	Conn *nats.Conn
}

type Config struct {
	Host string
	Port int
}

func NewNatsClient(conf Config) *NatsClient {
	conn, err := nats.Connect(fmt.Sprintf("nats://%s:%d", conf.Host, conf.Port))
	if err != nil {
		panic(fmt.Sprintf("ошибка подключения к NATS: %s", err))
	}

	return &NatsClient{Conn: conn}
}

func (n *NatsClient) Publish(subject string, data []byte) error {
	return n.Conn.Publish(subject, data)
}

func (n *NatsClient) Subscribe(subject string, handler func(msg *nats.Msg)) (*nats.Subscription, error) {
	return n.Conn.Subscribe(subject, handler)
}

func (n *NatsClient) Close() {
	n.Conn.Close()
}
