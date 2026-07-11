package json

import (
	"encoding/json"
	"os"

	"github.com/binarysoupdev/go-commando/errors"
)

// Marshal encodes the given typed object as JSON.
func Marshal[T any](obj T) ([]byte, error) {
	return json.Marshal(obj)
}

// Marshal encodes the given typed object to a JSON file.
func MarshalFile[T any](obj T, path string) error {
	return MarshalFilePretty(obj, path, "")
}

// Marshal encodes the given typed object to a JSON file, indenting each line.
func MarshalFilePretty[T any](obj T, path string, indent string) error {
	bytes, err := json.MarshalIndent(obj, "", indent)
	if err != nil {
		return errors.Chain(err, "error marshaling JSON")
	}

	err = os.WriteFile(path, bytes, 0666)
	if err != nil {
		return errors.Chain(err, "error writing JSON file")
	}

	return nil
}
