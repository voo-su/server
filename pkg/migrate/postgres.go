package migrate

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
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
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("не удалось создать экземпляр базы данных: %v", err)
	}

	migrator, err := migrate.NewWithInstance("migration_embeded_sql_files", m.srcDriver, "psql_db", driver)
	if err != nil {
		return fmt.Errorf("не удалось создать миграцию: %v", err)
	}

	defer func() {
		migrator.Close()
	}()

	if err = migrator.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("не удалось применить миграции: %v", err)
	}

	return nil
}
