package config

type Nats struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
