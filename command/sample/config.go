package sample

import (
	"fmt"

	"github.com/binarysoupdev/go-commando/command"
	"github.com/binarysoupdev/go-commando/errors"
	"github.com/binarysoupdev/go-commando/json"
)

type Config struct {
	Data string
}

type ConfigCommand struct {
	command.CommandBase

	ConfigLoader json.Loader[Config]
	Config       Config
}

func NewConfigCommand(configLoader json.Loader[Config]) *ConfigCommand {
	return &ConfigCommand{
		CommandBase:  command.NewCommandBase("config", "load and display the config file"),
		ConfigLoader: configLoader,
	}
}

func (cmd *ConfigCommand) Initialize() error {
	var err error

	cmd.Config, err = cmd.ConfigLoader.Load()
	if err != nil {
		return errors.Chain(err, "error loading config")
	}
	return nil
}

func (cmd ConfigCommand) Run(args []string) error {
	fmt.Printf("Data: %s\n", cmd.Config.Data)
	return nil
}
