package command

import (
	"github.com/binarysoupdev/go-extensions/errors"
	"github.com/binarysoupdev/go-extensions/json"
)

// Command component for loading a config object from JSON.
type ConfigCommand[Config any] struct {
	ConfigLoader json.Loader[Config]
	Config       Config
}

func NewConfigCommand[Config any](loader json.Loader[Config]) ConfigCommand[Config] {
	return ConfigCommand[Config]{
		ConfigLoader: loader,
	}
}

// Load the config from the JSON file and set cmd.Config.
func (cmd *ConfigCommand[Config]) LoadConfig() error {
	cfg, err := cmd.ConfigLoader.Load()
	if err != nil {
		return errors.Chain(err, "error loading config")
	}

	cmd.Config = cfg
	return nil
}
