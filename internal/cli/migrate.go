package cli

import (
	"gorm.io/gorm"
	"voo.su/pkg/migrate"

	"github.com/urfave/cli/v2"
	"voo.su/internal/config"
)

type MigrateProvider struct {
	Conf *config.Config
	DB   *gorm.DB
}

func Migrate(ctx *cli.Context, app *MigrateProvider) error {
	if err := migrate.Postgres(app.Conf); err != nil {
		return err
	}

	return nil
}
