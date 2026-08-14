package logger

import (
	"log"
)

var logger *log.Logger

// Set the package level logger.
func SetLogger(log *log.Logger) {
	logger = log
}
