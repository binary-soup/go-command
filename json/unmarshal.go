package json

import (
	"encoding/json"
	"os"

	"github.com/binarysoupdev/go-commando/errors"
)

// Unmarshal decodes the given JSON bytes into the given type.
func Unmarshal[T any](bytes []byte) (T, error) {
	var obj T
	return obj, json.Unmarshal(bytes, &obj)
}

// Unmarshal decodes the given JSON file into the given type.
func UnmarshalFile[T any](path string) (T, error) {
	var obj T

	bytes, err := os.ReadFile(path)
	if err != nil {
		return obj, errors.Chain(err, "error reading JSON file")
	}

	obj, err = Unmarshal[T](bytes)
	if err != nil {
		return obj, errors.Chain(err, "error unmarshaling JSON")
	}

	return obj, nil
}
