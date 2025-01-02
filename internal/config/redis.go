// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package config

import (
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Auth     string `yaml:"auth"`
	Database int    `yaml:"database"`
}

func (r *Redis) Options() *redis.Options {
	return &redis.Options{
		Addr:        r.Host,
		Password:    r.Auth,
		DB:          r.Database,
		ReadTimeout: -1,
	}
}
