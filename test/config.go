package test

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Grpc struct {
	Token string `yaml:"token"`
}

type Config struct {
	Grpc *Grpc `yaml:"grpc"`
}

func GetConfig() *Config {
	_, file, _, _ := runtime.Caller(0)
	content, err := os.ReadFile(filepath.Join([]string{filepath.Dir(filepath.Dir(file)), "./config-dev/voo-su-test.yaml"}...))
	if err != nil {
		panic(err)
	}

	var conf Config
	if err := yaml.Unmarshal(content, &conf); err != nil {
		log.Println(err)
		panic(fmt.Sprintf("Ошибка при разборе: %v", err))
	}

	return &conf
}
