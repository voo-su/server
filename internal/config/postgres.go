// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package config

import (
	"fmt"
)

type Postgres struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func (p *Postgres) GetDsn() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Europe/Moscow",
		p.Host,
		p.Port,
		p.Username,
		p.Password,
		p.Database,
	)
}
