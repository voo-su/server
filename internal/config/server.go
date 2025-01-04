package config

import "fmt"

type Server struct {
	Http      Http      `yaml:"http"`
	Websocket Websocket `yaml:"ws"`
	Tcp       Tcp       `yaml:"tcp"`
	Grpc      Grpc      `yaml:"grpc"`
}

type Http struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func (h *Http) GetHttp() string {
	return fmt.Sprintf("%s:%d", h.Host, h.Port)
}

type Websocket struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func (w *Websocket) GetWebsocket() string {
	return fmt.Sprintf("%s:%d", w.Host, w.Port)
}

type Tcp struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func (t *Tcp) GetTcp() string {
	return fmt.Sprintf("%s:%d", t.Host, t.Port)
}

type Grpc struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Protocol string `yaml:"protocol"`
}

func (g *Grpc) GetGrpc() string {
	return fmt.Sprintf("%s:%d", g.Host, g.Port)
}
