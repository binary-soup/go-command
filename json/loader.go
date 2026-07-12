package json

import (
	"os"

	"github.com/binarysoupdev/go-commando/errors"
)

type Loader[T any] struct {
	Path string
}

func NewLoader[T any](path string) Loader[T] {
	return Loader[T]{
		Path: path,
	}
}

func (u Loader[T]) ValidatePath() error {
	_, err := os.Stat(u.Path)
	if err != nil {
		return errors.Format("path \"%s\" not found/accessible", u.Path)
	}
	return nil
}

func (u Loader[T]) Load() (T, error) {
	return UnmarshalFile[T](u.Path)
}
