package config

import (
	"github.com/binarysoupdev/go-commando/errors"
	"github.com/binarysoupdev/go-commando/json"
)

type Version struct {
	Version int `json:"version"`
}

func (l *Loader[T]) LoadVersion() (int, error) {
	v, err := json.UnmarshalFile[Version](l.ConfigPath)
	if err != nil {
		return -1, errors.Chain(err, "error loading version JSON")
	}

	return v.Version, nil
}
