package config

import (
	"github.com/binarysoupdev/go-commando/errors"
	"github.com/binarysoupdev/go-commando/json"
)

type Version struct {
	Version int `json:"version"`
}

func (l *Loader[T]) LoadVersion() error {
	v, err := json.UnmarshalFile[Version](l.ConfigPath)
	if err != nil {
		return errors.Chain(err, "error loading version JSON")
	}

	l.ConfigVersion = v.Version
	return nil
}
