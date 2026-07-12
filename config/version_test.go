package config_test

import (
	"testing"

	"github.com/binarysoupdev/go-commando/config"
	"github.com/binarysoupdev/go-commando/json"
	"github.com/binarysoupdev/tinsel/file"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadVersionReturnsErrorWhenUnmarshalJsonFileFails(t *testing.T) {
	//-- arrange
	loader := config.NewLoader[config.Version]("invalid")

	//-- act
	_, res := loader.LoadVersion()

	//-- assert
	require.ErrorContains(t, res, "error loading version JSON")
}

func TestLoadVersionReturnsVersionAndNoErrorWhenValid(t *testing.T) {
	//-- arrange
	loader := config.NewLoader[config.Version](file.NewPath(t, "config.json"))

	CONFIG := config.Version{
		Version: 1,
	}

	err := json.MarshalFile(CONFIG, loader.ConfigPath)
	require.NoError(t, err)

	//-- act
	res, err := loader.LoadVersion()

	//-- assert
	require.NoError(t, err)
	assert.Equal(t, CONFIG.Version, res)
}
