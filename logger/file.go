package logger

import (
	"log"
	"os"

	"github.com/binarysoupdev/go-extensions/errors"
)

// Create/open a log file at the given path and configure as the logger output.
// Returns the file and any errors.
//
// The calling code is responsible for closing the file. The file is write only.
func OpenFile(path string) (*os.File, error) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, errors.Chain(err, "error opening/creating log file")
	}

	SetLogger(log.New(file, "", log.Ldate|log.Ltime))
	return file, nil
}
