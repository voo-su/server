package migrate

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"log"
	vooSu "voo.su"
	"voo.su/internal/config"

	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

type Migrator struct {
	srcDriver source.Driver
}

func MustGetNewMigrator(sqlFiles embed.FS, dirName string) *Migrator {
	d, err := iofs.New(sqlFiles, dirName)
	if err != nil {
		panic(err)
	}

	return &Migrator{
		srcDriver: d,
	}
}

func (m *Migrator) ApplyMigrations(db *sql.DB) error {
	// Создание драйвера PostgreSQL для указанной базы данных.
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("не удалось создать экземпляр базы данных: %v", err)
	}

	// Создание нового мигратора с использованием встроенных SQL-файлов.
	migrator, err := migrate.NewWithInstance("migration_embeded_sql_files", m.srcDriver, "psql_db", driver)
	if err != nil {
		return fmt.Errorf("не удалось создать миграцию: %v", err)
	}

	defer func() {
		// Закрытие мигратора по завершении работы.
		migrator.Close()
	}()

	// Применение миграций к базе данных.
	if err = migrator.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("не удалось применить миграции: %v", err)
	}

	return nil
}

func Postgres(conf *config.Config) error {
	conn, err := sql.Open("postgres", conf.Postgresql.GetDsn())
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	defer conn.Close()

	// Создание нового экземпляра Migrator с использованием встроенных SQL-файлов миграций.
	migrator := MustGetNewMigrator(vooSu.Migration(), "migrations")

	// Применение миграций к базе данных.
	if err = migrator.ApplyMigrations(conn); err != nil {
		log.Fatalf("Ошибка при применении миграций: %v", err)
	}

	return nil
}
