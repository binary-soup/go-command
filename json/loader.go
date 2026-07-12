package json

import (
	"os"

	"github.com/binarysoupdev/go-commando/errors"
)

// Loader caches a file path so the JSON file can be loaded only once requested.
type Loader[T any] struct {
	Path string
}

func NewLoader[T any](path string) Loader[T] {
	return Loader[T]{
		Path: path,
	}
}

// Validate the path points to a real a file.
// Does NOT check if the file is valid JSON.
func (u Loader[T]) ValidatePath() error {
	_, err := os.Stat(u.Path)
	if err != nil {
		return errors.Format("path \"%s\" not found/accessible", u.Path)
	}
	return nil
}

// Load the JSON file at the cached path.
func (u Loader[T]) Load() (T, error) {
	return UnmarshalFile[T](u.Path)
}
