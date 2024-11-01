package provider

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
	app := &cli.App{
		Name: "Voo.su",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Value:       "/etc/voo-su/",
				Usage:       "Путь к папке с файлами конфигурации",
				DefaultText: "/etc/voo-su/",
			},
		},
	}
	return &App{app: app}
}

func (c *App) Register(cm Command) {
	c.app.Commands = append(c.app.Commands, c.createCommand(cm))
}

func (c *App) createCommand(cm Command) *cli.Command {
	cd := &cli.Command{
		Name:        cm.Name,
		Usage:       cm.Usage,
		Flags:       cm.Flags,
		Subcommands: c.createSubcommands(cm.Subcommands),
	}

	if cm.Action != nil {
		cd.Action = func(ctx *cli.Context) error {
			conf := config.LoadConfig(ctx.String("config"))
			return cm.Action(ctx, conf)
		}
	}

	return cd
}

func (c *App) createSubcommands(commands []Command) []*cli.Command {
	var subcommands []*cli.Command
	for _, subCmd := range commands {
		subcommands = append(subcommands, c.createCommand(subCmd))
	}
	return subcommands
}

func (c *App) Run() {
	_ = c.app.Run(os.Args)
}
