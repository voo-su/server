// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"time"
	"voo.su/pkg/encrypt"
	"voo.su/pkg/strutil"
)

type Config struct {
	sid        string
	App        *App        `yaml:"app"`
	Server     *Server     `yaml:"server"`
	Postgres   *Postgres   `yaml:"postgres"`
	ClickHouse *ClickHouse `yaml:"clickhouse"`
	Minio      *Minio      `yaml:"minio"`
	Redis      *Redis      `yaml:"redis"`
	Nats       *Nats       `yaml:"nats"`
	Email      *Email      `yaml:"email"`
	Push       *Push       `yaml:"push"`
}

func New(filename string) *Config {
	//loc, _ := time.LoadLocation("Europe/Moscow")
	//time.Local = loc

	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var conf Config
	if err := yaml.Unmarshal(content, &conf); err != nil {
		fmt.Println(err)
		panic(fmt.Sprintf("Ошибка при разборе: %v", err))
	}

	conf.sid = encrypt.Md5(fmt.Sprintf("%d%s", time.Now().UnixNano(), strutil.Random(6)))

	return &conf
}

func (c *Config) ServerId() string {
	return c.sid
}
