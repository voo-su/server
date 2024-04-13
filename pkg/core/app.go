package core

import (
	"github.com/urfave/cli/v2"
	"os"
	"voo.su/internal/config"
)

type App struct {
	app *cli.App
}

type Action func(ctx *cli.Context, conf *config.Config) error

type Command struct {
	Name        string
	Usage       string
	Flags       []cli.Flag
	Action      Action
	Subcommands []Command
}

func NewApp() *App {
	return &App{
		app: &cli.App{
			Name: "Voo.su",
		},
	}
}

func (c *App) Register(cm Command) {
	c.app.Commands = append(c.app.Commands, c.command(cm))
}

func (c *App) command(cm Command) *cli.Command {
	cd := &cli.Command{
		Name:  cm.Name,
		Usage: cm.Usage,
		Flags: make([]cli.Flag, 0),
	}

	if len(cm.Subcommands) > 0 {
		for _, v := range cm.Subcommands {
			cd.Subcommands = append(cd.Subcommands, c.command(v))
		}
	} else {
		if cm.Flags != nil && len(cm.Flags) > 0 {
			cd.Flags = append(cd.Flags, cm.Flags...)
		}

		var isConfig bool

		for _, flag := range cd.Flags {
			if flag.Names()[0] == "config" {
				isConfig = true
				break
			}
		}

		if !isConfig {
			cd.Flags = append(cd.Flags, &cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Value:       "/etc/voo-su/voo-su.yaml",
				Usage:       "Путь к файлу конфигурации",
				DefaultText: "/etc/voo-su/voo-su.yaml",
			})
		}

		if cm.Action != nil {
			cd.Action = func(ctx *cli.Context) error {
				return cm.Action(ctx, config.New(ctx.String("config")))
			}
		}
	}

	return cd
}

func (c *App) Run() {
	_ = c.app.Run(os.Args)
}
