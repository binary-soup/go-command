package types_test

import (
	"testing"

	"github.com/binarysoupdev/go-commando/types"
	"github.com/stretchr/testify/require"
)

func TestIsVersionUnsupportedReturnsTrueWhenVersionLessThanMin(t *testing.T) {
	//-- arrange
	const CURRENT_VERSION = 1
	version := types.Version(0)

	//-- act
	res := version.IsUnsupported(CURRENT_VERSION)

	//-- assert
	require.True(t, res)
}

func TestIsVersionUnsupportedReturnsTrueWhenVersionGreaterThanCurrent(t *testing.T) {
	//-- arrange
	const CURRENT_VERSION = 1
	version := types.Version(CURRENT_VERSION + 1)

	//-- act
	res := version.IsUnsupported(CURRENT_VERSION)

	//-- assert
	require.True(t, res)
}

func TestIsVersionUnsupportedReturnsFalseWhenVersionValid(t *testing.T) {
	//-- arrange
	const CURRENT_VERSION = 1
	version := types.Version(CURRENT_VERSION)

	//-- act
	res := version.IsUnsupported(CURRENT_VERSION)

	//-- assert
	require.False(t, res)
}

func TestIsVersionOutOfDateReturnsTrueWhenVersionLessThanCurrent(t *testing.T) {
	//-- arrange
	const CURRENT_VERSION = 1
	version := types.Version(CURRENT_VERSION - 1)

	//-- act
	res := version.IsOutOfDate(CURRENT_VERSION)

	//-- assert
	require.True(t, res)
}

func TestIsVersionOutOfDateReturnsFalseWhenVersionEqualsCurrent(t *testing.T) {
	//-- arrange
	const CURRENT_VERSION = 1
	version := types.Version(CURRENT_VERSION)

	//-- act
	res := version.IsOutOfDate(CURRENT_VERSION)

	//-- assert
	require.False(t, res)
}
