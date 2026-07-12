package config_test

import (
	"fmt"
	"testing"

	"github.com/binarysoupdev/go-commando/config"
	"github.com/binarysoupdev/go-commando/json"
	"github.com/binarysoupdev/tinsel/file"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestConfig struct {
	Data string
}

func TestValidatePathReturnsErrorWhenPathNotFound(t *testing.T) {
	//-- arrange
	loader := config.NewLoader[TestConfig]("invalid")

	//-- act
	res := loader.ValidatePath()

	//-- assert
	require.ErrorContains(t, res, fmt.Sprintf("path \"%s\" not found/accessible", loader.ConfigPath))
}

func TestValidatePathReturnsNoErrorWhenPathValid(t *testing.T) {
	//-- arrange
	loader := config.NewLoader[TestConfig](file.NewPath(t, ""))

	//-- act
	res := loader.ValidatePath()

	//-- assert
	require.NoError(t, res)
}

func TestLoadConfigReturnsErrorWhenUnmarshalJsonFileFails(t *testing.T) {
	//-- arrange
	loader := config.NewLoader[TestConfig]("invalid")

	//-- act
	res := loader.LoadConfig()

	//-- assert
	require.ErrorContains(t, res, "error loading config JSON")
}

func TestLoadConfigReturnsNoErrorAndLoadsConfigWhenValid(t *testing.T) {
	//-- arrange
	loader := config.NewLoader[TestConfig](file.NewPath(t, "config.json"))

	CONFIG := TestConfig{
		Data: "foobar",
	}

	err := json.MarshalFile(CONFIG, loader.ConfigPath)
	require.NoError(t, err)

	//-- act
	res := loader.LoadConfig()

	//-- assert
	require.NoError(t, res)
	assert.Equal(t, CONFIG, loader.Config)
}
