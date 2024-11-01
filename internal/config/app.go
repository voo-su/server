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
	Env string `yaml:"env"`
	Log string `yaml:"log"`
}

type Config struct {
	sid        string
	App        *App        `yaml:"app"`
	Server     *Server     `yaml:"server"`
	Redis      *Redis      `yaml:"redis"`
	Postgresql *Postgresql `yaml:"postgresql"`
	Jwt        *Jwt        `yaml:"jwt"`
	Cors       *Cors       `yaml:"cors"`
	File       *File       `yaml:"file"`
	Email      *Email      `yaml:"email"`
}

func loadFile(filename string, target interface{}) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("ошибка чтения файла: %v", err)
	}

	if err := yaml.Unmarshal(content, target); err != nil {
		return fmt.Errorf("ошибка разбора YAML: %v", err)
	}

	return nil
}

// LoadConfig загружает конфигурацию из всех YAML файлов
func LoadConfig(configPath string) *Config {
	conf := &Config{}
	filenames := []string{
		configPath + "app.yaml",
		configPath + "server.yaml",
		configPath + "database.yaml",
		configPath + "auth.yaml",
		configPath + "file.yaml",
		configPath + "email.yaml",
	}

	for _, filename := range filenames {
		if err := loadFile(filename, conf); err != nil {
			fmt.Printf("Предупреждение: Не удалось загрузить конфигурацию из %s: %v\n", filename, err)
		}
	}

	conf.sid = encrypt.Md5(fmt.Sprintf("%d%s", time.Now().UnixNano(), strutil.Random(6)))
	return conf
}

func (c *Config) ServerId() string {
	return c.sid
}
