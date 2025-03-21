package provider

import (
	"github.com/urfave/cli/v2"
	"os"
	"voo.su/internal/config"
	"voo.su/pkg/logger"
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
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "config",
					Aliases:     []string{"c"},
					Value:       "/etc/voo-su/voo-su.yaml",
					DefaultText: "/etc/voo-su/voo-su.yaml",
				},
			},
		},
	}
}

func (a *App) Register(cm Command) {
	a.app.Commands = append(a.app.Commands, a.createCommand(cm))
}

func (a *App) createCommand(cm Command) *cli.Command {
	cd := &cli.Command{
		Name:        cm.Name,
		Usage:       cm.Usage,
		Flags:       cm.Flags,
		Subcommands: a.createSubcommands(cm.Subcommands),
	}

	if cm.Action != nil {
		cd.Action = func(ctx *cli.Context) error {
			return cm.Action(ctx, config.New(ctx.String("config")))
		}
	}

	return cd
}

func (a *App) createSubcommands(commands []Command) []*cli.Command {
	var subcommands []*cli.Command
	for _, subCmd := range commands {
		subcommands = append(subcommands, a.createCommand(subCmd))
	}
	return subcommands
}

func (a *App) Run() {
	logger.InitLogger(logger.LevelInfo, a.app.Name)
	_ = a.app.Run(os.Args)
}
