package infrastructure

import (
	clickHouseDriver "github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Source struct {
	PG  *gorm.DB
	CH  *clickHouseDriver.Conn
	Rds *redis.Client
}

func NewSource(
	postgres *gorm.DB,
	clickHouse *clickHouseDriver.Conn,
	redis *redis.Client,
) *Source {
	return &Source{
		PG:  postgres,
		CH:  clickHouse,
		Rds: redis,
	}
}

func (s *Source) Postgres() *gorm.DB {
	return s.PG
}

func (s *Source) Redis() *redis.Client {
	return s.Rds
}

func (s *Source) ClickHouse() *clickHouseDriver.Conn {
	return s.CH
}
