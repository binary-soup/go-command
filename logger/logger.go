package logger

import (
	"log"
)

var logger *log.Logger

// Get the package level logger. Can be used to further configure it.
func GetLogger() *log.Logger {
	return logger
}

// Set the package level logger.
func SetLogger(log *log.Logger) {
	logger = log
}
