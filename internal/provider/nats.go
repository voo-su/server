// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package provider

import (
	"voo.su/internal/config"
	"voo.su/pkg/nats"
)

type Nats struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func NewNatsClient(conf *config.Config) nats.INatsClient {
	return nats.NewNatsClient(nats.Config{
		Host: conf.Nats.Host,
		Port: conf.Nats.Port,
	})
}
