package sample_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/binarysoupdev/go-commando/command/sample"
	"github.com/binarysoupdev/go-commando/test"
	"github.com/binarysoupdev/go-extensions/json"
	"github.com/binarysoupdev/tinsel/pipe"
	"github.com/stretchr/testify/suite"
)

type ConfigTestSuite struct {
	test.CommandSuite[*sample.ConfigCommand]
	ConfigLoader json.Loader[sample.Config]
	Config       sample.Config
}

func TestConfigCommandSuite(t *testing.T) {
	s := ConfigTestSuite{
		ConfigLoader: json.NewLoader[sample.Config](filepath.Join(t.TempDir(), "config.json")),
	}

	s.CommandSuite = test.NewCommandSuite(sample.NewConfigCommand(s.ConfigLoader))
	suite.Run(t, &s)
}

func (s *ConfigTestSuite) SetupTest() {
	s.Config = sample.Config{
		Version: sample.CONFIG_VERSION,
		Data:    "foobar",
	}

	err := json.MarshalFile(s.Config, s.ConfigLoader.Path)
	s.Require().NoError(err)
}

//=======================================================

func (s *ConfigTestSuite) TestRunFailsWhenLoadConfigFails() {
	//-- arrange
	err := os.Remove(s.ConfigLoader.Path)
	s.Require().NoError(err)

	//-- act
	s.RunCommand()

	//-- assert
	s.RequireResultFail("error loading config")
}

func (s *ConfigTestSuite) TestRunFailsWhenConfigVersionIsLessThanOne() {
	//-- arrange
	s.Config.Version = 0
	err := json.MarshalFile(s.Config, s.ConfigLoader.Path)
	s.Require().NoError(err)

	//-- act
	s.RunCommand()

	//-- assert
	s.RequireResultFail(fmt.Sprintf("config version \"%d\" unsupported", s.Config.Version))
}

func (s *ConfigTestSuite) TestRunFailsWhenConfigVersionIsGreaterThanCurrent() {
	//-- arrange
	s.Config.Version = sample.CONFIG_VERSION + 1
	err := json.MarshalFile(s.Config, s.ConfigLoader.Path)
	s.Require().NoError(err)

	//-- act
	s.RunCommand()

	//-- assert
	s.RequireResultFail(fmt.Sprintf("config version \"%d\" unsupported", s.Config.Version))
}

func (s *ConfigTestSuite) TestRunFailsWhenConfigVersionIsOutOfDate() {
	//-- arrange
	s.Config.Version = sample.CONFIG_VERSION - 1
	err := json.MarshalFile(s.Config, s.ConfigLoader.Path)
	s.Require().NoError(err)

	//-- act
	s.RunCommand()

	//-- assert
	s.RequireResultFail(fmt.Sprintf("config version \"%d\" out-of-date", s.Config.Version))
}

func (s *ConfigTestSuite) TestRunPassesAndPrintsConfig() {
	//-- arrange
	out := pipe.OpenStdout(2)
	defer out.Close()

	//-- act
	s.RunCommand()

	//-- assert
	s.RequireResultPass()
	s.Assert().Contains(out.ReadLine(), fmt.Sprintf("Version: %d", s.Config.Version))
	s.Assert().Contains(out.ReadLine(), fmt.Sprintf("Data: %s", s.Config.Data))
}
