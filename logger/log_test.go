package logger_test

import (
	"bytes"
	"errors"
	"log"
	"testing"

	"github.com/binarysoupdev/go-commando/logger"
	"github.com/stretchr/testify/assert"
)

func TestLogLogsMessage(t *testing.T) {
	//-- arrange
	buffer := &bytes.Buffer{}
	const MESSAGE = "message"

	logger.SetLogger(log.New(buffer, "", 0))

	//-- act
	logger.Log(MESSAGE)

	//-- assert
	assert.Contains(t, buffer.String(), MESSAGE)
}

func TestLogfLogsFormattedMessage(t *testing.T) {
	//-- arrange
	buffer := &bytes.Buffer{}
	const MESSAGE = "message"

	logger.SetLogger(log.New(buffer, "", 0))

	//-- act
	logger.Logf("%s", MESSAGE)

	//-- assert
	assert.Contains(t, buffer.String(), MESSAGE)
}

func TestLogErrorLogsError(t *testing.T) {
	//-- arrange
	buffer := &bytes.Buffer{}
	const MESSAGE = "message"

	logger.SetLogger(log.New(buffer, "", 0))

	//-- act
	logger.LogError(errors.New(MESSAGE))

	//-- assert
	assert.Contains(t, buffer.String(), MESSAGE)
}
