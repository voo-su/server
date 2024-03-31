package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"voo.su/internal/config"
	"voo.su/pkg/logger"

	"github.com/urfave/cli/v2"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	app := Initialize(ctx, config.New(parseConfigArg()))
	logger.InitLogger(fmt.Sprintf("%s/cmd.log", app.Config.LogPath()), logger.LevelInfo, "cmd")
	newApp(ctx, app.Commands.SubCommands())
}

func newApp(ctx context.Context, commands []*cli.Command) {
	cmd := cli.NewApp()
	cmd.Name = "Voo.su"
	cmd.Usage = "Интерфейс командной строки"
	cmd.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "config",
			Aliases:     []string{"c"},
			Value:       "./init/voo-su.yaml",
			Usage:       "Путь к файлу конфигурации",
			DefaultText: "./init/voo-su.yaml",
		},
	}
	cmd.Commands = commands
	if err := cmd.RunContext(ctx, os.Args); err != nil {
		logger.Std().Error(fmt.Sprintf("Ошибка выполнения команды: %s", err.Error()))
	}
}

func parseConfigArg() string {
	var conf string
	flag.StringVar(&conf, "config", "./init/voo-su.yaml", "Путь к файлу конфигурации")
	flag.Parse()
	return conf
}
