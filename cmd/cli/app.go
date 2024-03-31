package main

import (
	"voo.su/internal/config"
	"voo.su/internal/transport/cli/command"
)

type AppProvider struct {
	Config   *config.Config
	Commands *command.Commands
}
