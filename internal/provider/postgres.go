package provider

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
	"voo.su/internal/config"
)

func NewPostgresqlClient(conf *config.Config) *gorm.DB {
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	if conf.App.Env == "dev" {
		writer, _ := os.OpenFile(conf.App.LogPath("postgres.log"), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
		gormConfig.Logger = logger.New(
			log.New(writer, "", log.LstdFlags),
			logger.Config{
				SlowThreshold:             200 * time.Millisecond,
				LogLevel:                  logger.Warn,
				IgnoreRecordNotFoundError: true,
			},
		)
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: conf.Postgres.GetDsn(),
	}), gormConfig)
	if err != nil {
		panic(fmt.Errorf("ошибка подключения к postgres: %v", err))
	}

	if db.Error != nil {
		panic(fmt.Errorf("ошибка базы данных: %v", err))
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
