package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
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
	Manager    *Manager    `yaml:"manager"`
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
		log.Println(err)
		panic(fmt.Sprintf("Ошибка при разборе: %v", err))
	}

	conf.sid = encrypt.Md5(fmt.Sprintf("%d%s", time.Now().UnixNano(), strutil.Random(6)))

	return &conf
}

func (c *Config) ServerId() string {
	return c.sid
}
