package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"time"
	"voo.su/pkg/encrypt"
	"voo.su/pkg/strutil"
)

type App struct {
	Http      int    `yaml:"http"`
	Websocket int    `yaml:"ws"`
	Tcp       int    `yaml:"tcp"`
	Env       string `yaml:"env"`
	Log       string `yaml:"log"`
}

type Config struct {
	sid        string
	App        *App        `yaml:"app"`
	Redis      *Redis      `yaml:"redis"`
	Postgresql *Postgresql `yaml:"postgresql"`
	Jwt        *Jwt        `yaml:"jwt"`
	Cors       *Cors       `yaml:"cors"`
	File       *File       `yaml:"file"`
	Email      *Email      `yaml:"email"`
	Nsq        *Nsq        `yaml:"nsq"`
}

func New(filename string) *Config {
	//loc, _ := time.LoadLocation("Europe/Moscow")
	//time.Local = loc

	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var conf Config
	if yaml.Unmarshal(content, &conf) != nil {
		panic(fmt.Sprintf("Ошибка при разборе: %v", err))
	}

	conf.sid = encrypt.Md5(fmt.Sprintf("%d%s", time.Now().UnixNano(), strutil.Random(6)))

	return &conf
}

func (c *Config) ServerId() string {
	return c.sid
}
