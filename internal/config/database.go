package config

import "fmt"

type Postgresql struct {
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
