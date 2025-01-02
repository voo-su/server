// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package infrastructure

import (
	clickHouseDriver "github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Source struct {
	DB         *gorm.DB
	clickHouse clickHouseDriver.Conn
	redis      *redis.Client
}

func NewSource(
	db *gorm.DB,
	clickHouse clickHouseDriver.Conn,
	redis *redis.Client,
) *Source {
	return &Source{
		DB:         db,
		clickHouse: clickHouse,
		redis:      redis,
	}
}

func (s *Source) Postgres() *gorm.DB {
	return s.DB
}

func (s *Source) Redis() *redis.Client {
	return s.redis
}

func (s *Source) ClickHouse() clickHouseDriver.Conn {
	return s.clickHouse
}
