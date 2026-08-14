package logger

import (
	"log"
	"os"

	"github.com/binarysoupdev/go-extensions/errors"
)

// Create/open a log file at the given path and configure as the logger output.
// Returns the logger, file, and any errors.
func OpenFile(path string) (*log.Logger, *os.File, error) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, nil, errors.Chain(err, "error opening/creating log file")
	}

	logger = log.New(file, "", log.Ldate|log.Ltime)
	return logger, file, nil
}
