package logger

import (
	"log"
)

var logger *log.Logger

func SetLogger(log *log.Logger) {
	logger = log
}
