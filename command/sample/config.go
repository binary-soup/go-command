package sample

import (
	"fmt"

	"github.com/binarysoupdev/go-commando/command"
	"github.com/binarysoupdev/go-commando/errors"
	"github.com/binarysoupdev/go-commando/json"
	"github.com/binarysoupdev/go-commando/types"
)

const CONFIG_VERSION = 2

type Config struct {
	Version types.Version `json:"version"`
	Data    string        `json:"data"`
}

type ConfigCommand struct {
	command.CommandBase
	command.ConfigCommand[Config]
}

func NewConfigCommand(configLoader json.Loader[Config]) *ConfigCommand {
	return &ConfigCommand{
		CommandBase:   command.NewCommandBase("config", "load and display the config file"),
		ConfigCommand: command.NewConfigCommand(configLoader),
	}
}

func (cmd *ConfigCommand) Initialize() error {
	return cmd.LoadConfig()
}

func (cmd ConfigCommand) Run(args []string) error {
	if cmd.Config.Version.IsUnsupported(CONFIG_VERSION) {
		return errors.Format("version \"%d\" unsupported", cmd.Config.Version)
	}

	if cmd.Config.Version.IsOutOfDate(CONFIG_VERSION) {
		return errors.Format("config version \"%d\" out-of-date", cmd.Config.Version)
	}

	fmt.Printf("Version: %d\nData: %s\n", cmd.Config.Version, cmd.Config.Data)
	return nil
}
