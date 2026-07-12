package json_test

import (
	"fmt"
	"testing"

	"github.com/binarysoupdev/go-commando/json"
	"github.com/binarysoupdev/tinsel/file"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestObject struct {
	Data string `json:"data"`
}

func TestLoaderValidatePathReturnsErrorWhenPathNotFound(t *testing.T) {
	//-- arrange
	loader := json.NewLoader[TestObject]("invalid")

	//-- act
	res := loader.ValidatePath()

	//-- assert
	require.ErrorContains(t, res, fmt.Sprintf("path \"%s\" not found/accessible", loader.Path))
}

func TestLoaderValidatePathReturnsNoErrorWhenPathValid(t *testing.T) {
	//-- arrange
	loader := json.NewLoader[TestObject](file.NewPath(t, ""))

	//-- act
	res := loader.ValidatePath()

	//-- assert
	require.NoError(t, res)
}

func TestLoaderUnmarshalReturnsErrorWhenUnmarshalJsonFileFails(t *testing.T) {
	//-- arrange
	loader := json.NewLoader[TestObject]("invalid")

	//-- act
	_, res := loader.Load()

	//-- assert
	require.ErrorContains(t, res, "error reading JSON file")
}

func TestLoaderUnmarshalReturnsObjectAndNoWhenValid(t *testing.T) {
	//-- arrange
	loader := json.NewLoader[TestObject](file.NewPath(t, "data.json"))

	OBJECT := TestObject{
		Data: "foobar",
	}

	err := json.MarshalFile(OBJECT, loader.Path)
	require.NoError(t, err)

	//-- act
	res, err := loader.Load()

	//-- assert
	require.NoError(t, err)
	assert.Equal(t, OBJECT, res)
}
