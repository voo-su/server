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
	Redis      *Redis      `yaml:"redis"`
	Postgres   *Postgres   `yaml:"postgres"`
	ClickHouse *ClickHouse `yaml:"clickhouse"`
	Minio      *Minio      `yaml:"minio"`
	Email      *Email      `yaml:"email"`
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
