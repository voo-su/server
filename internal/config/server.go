package config

type Server struct {
	Http      int `yaml:"http"`
	Websocket int `yaml:"ws"`
	Tcp       int `yaml:"tcp"`
}
