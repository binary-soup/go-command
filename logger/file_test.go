package logger_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/binarysoupdev/go-commando/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOpenFileReturnsErrorWhenInvalidPath(t *testing.T) {
	//-- act
	_, res := logger.OpenFile("invalid/path")

	//-- assert
	assert.ErrorContains(t, res, "error opening/creating log file")
}

func TestOpenFileReturnsFileAndNoErrorAndSetsLogger(t *testing.T) {
	//-- arrange
	PATH := filepath.Join(t.TempDir(), "log.txt")
	const MESSAGE = "message"

	//-- act
	file, err := logger.OpenFile(PATH)
	require.NoError(t, err)
	defer file.Close()

	//-- assert
	require.NotNil(t, logger.GetLogger())
	logger.Log(MESSAGE)

	bytes, err := os.ReadFile(PATH)
	require.NoError(t, err)
	assert.Contains(t, string(bytes), MESSAGE)
}
