package sample

import (
	"fmt"

	"github.com/binarysoupdev/go-commando/command"
	"github.com/binarysoupdev/go-extensions/errors"
	"github.com/binarysoupdev/go-extensions/json"
)

const CONFIG_VERSION = 2

type Config struct {
	Version int    `json:"version"`
	Data    string `json:"data"`
}

// Sample config command for loading, verifying, and displaying a config file.
type ConfigCommand struct {
	command.CommandBase
	command.ConfigCommand[Config]
}

func NewConfigCommand(configLoader json.Loader[Config]) *ConfigCommand {
	return &ConfigCommand{
		CommandBase:   command.NewCommandBase("config", "load, verify, and display a config file"),
		ConfigCommand: command.NewConfigCommand(configLoader),
	}
}

func (cmd *ConfigCommand) Initialize() error {
	return cmd.LoadConfig()
}

func (cmd ConfigCommand) Run(args []string) error {
	if cmd.Config.Version < 1 || cmd.Config.Version > CONFIG_VERSION {
		return errors.Format("config version \"%d\" unsupported", cmd.Config.Version)
	}

	if cmd.Config.Version < CONFIG_VERSION {
		return errors.Format("config version \"%d\" out-of-date", cmd.Config.Version)
	}

	fmt.Printf("Version: %d\nData: %s\n", cmd.Config.Version, cmd.Config.Data)
	return nil
}
