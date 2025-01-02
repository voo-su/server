// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package config

type Nats struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
