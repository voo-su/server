package config

import (
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
)

type Postgresql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type ClickHouse struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func (d *Postgresql) GetDsn() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Europe/Moscow",
		d.Host,
		d.Port,
		d.Username,
		d.Password,
		d.Database,
	)
}

func (d *ClickHouse) Options() *clickhouse.Options {
	return &clickhouse.Options{
		Addr: []string{fmt.Sprintf("%s:%d", d.Host, d.Port)},
		Auth: clickhouse.Auth{
			Database: d.Database,
			Username: d.Username,
			Password: d.Password,
		},
		ClientInfo: clickhouse.ClientInfo{
			Products: []struct {
				Name    string
				Version string
			}{
				{Name: "Voo.su", Version: ""},
			},
		},
		Debugf: func(format string, v ...interface{}) {
			fmt.Printf(format, v)
		},
		//TLS: &tls.Config{
		//	InsecureSkipVerify: true,
		//},
	}
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Auth     string `yaml:"auth"`
	Database int    `yaml:"database"`
}
