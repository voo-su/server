package cli

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"voo.su/internal/config"
	clickhouseRepo "voo.su/internal/infrastructure/clickhouse/repository"
	"voo.su/internal/provider"
	webPush "voo.su/pkg/push/web_push"
)

type GenerateProvider struct {
	Conf             *config.Config
	LoggerRepository *clickhouseRepo.LoggerRepository
}

func Generate(ctx *cli.Context, app *GenerateProvider) error {
	log.SetOutput(provider.NewLoggerWriter(app.Conf, os.Stdout, app.LoggerRepository))

	privateKey, publicKey, err := webPush.GenerateVAPIDKeys()
	if err != nil {
		log.Println(err)
	}

	fmt.Println(privateKey)
	fmt.Println(publicKey)

	return nil
}
