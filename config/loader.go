package config

import (
	"os"

	"github.com/binarysoupdev/go-commando/errors"
	"github.com/binarysoupdev/go-commando/json"
)

type Loader[Config any] struct {
	ConfigPath    string
	ConfigVersion int
	Config        Config
}

func NewLoader[Config any](path string) Loader[Config] {
	return Loader[Config]{
		ConfigPath: path,
	}
}

func (l Loader[Config]) ValidatePath() error {
	_, err := os.Stat(l.ConfigPath)
	if err != nil {
		return errors.Format("path \"%s\" not found/accessible", l.ConfigPath)
	}
	return nil
}

func (l *Loader[Config]) LoadConfig() error {
	var err error

	l.Config, err = json.UnmarshalFile[Config](l.ConfigPath)
	if err != nil {
		return errors.Chain(err, "error loading config JSON")
	}

	return err
}
