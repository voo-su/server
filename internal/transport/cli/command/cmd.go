package command

import (
	"github.com/urfave/cli/v2"
	"voo.su/internal/transport/cli/command/cron"
	"voo.su/pkg/cmdutil"
)

type Commands struct {
	CrontabCommand cron.Command
}

func (c *Commands) SubCommands() []*cli.Command {
	return cmdutil.ToSubCommand(c)
}
