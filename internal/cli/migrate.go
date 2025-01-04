package cli

import (
	"database/sql"
	"gorm.io/gorm"
	"log"
	vooSu "voo.su"
	"voo.su/pkg/migrate"

	"github.com/urfave/cli/v2"
	"voo.su/internal/config"
)

type MigrateProvider struct {
	Conf *config.Config
	DB   *gorm.DB
}

func Migrate(ctx *cli.Context, app *MigrateProvider) error {
	if err := Postgres(app.Conf); err != nil {
		return err
	}

	return nil
}

func Postgres(conf *config.Config) error {
	conn, err := sql.Open("postgres", conf.Postgres.GetDsn())
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
		return err
	}
	defer conn.Close()

	migrator := migrate.MustGetNewMigrator(vooSu.Migration(), "migrations")

	if err = migrator.ApplyMigrations(conn); err != nil {
		log.Fatalf("Ошибка при применении миграций: %v", err)
		return err
	}

	return nil
}
