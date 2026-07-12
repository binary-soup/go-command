package sample_test

import (
	"fmt"
	"testing"

	"github.com/binarysoupdev/go-commando/command/sample"
	"github.com/binarysoupdev/go-commando/test"
	"github.com/binarysoupdev/tinsel/pipe"
	"github.com/binarysoupdev/tinsel/rand"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type HelloTestSuite struct {
	test.CommandSuite[*sample.HelloCommand]
}

func TestHelloCommandSuite(t *testing.T) {
	suite.Run(t, &HelloTestSuite{
		CommandSuite: test.NewCommandSuite(sample.NewHelloCommand()),
	})
}

//=======================================================

func (s *HelloTestSuite) TestNameNotEmpty() {
	//-- act
	s.RunCommand("-name", "")

	//-- assert
	s.RequireResultFail("name cannot be empty")
}

func (s *HelloTestSuite) TestPrintName() {
	//-- arrange
	r := rand.New(42)
	var NAME = r.ASCII(10)

	out := pipe.OpenStdout(1)
	defer out.Close()

	//-- act
	s.RunCommand("-name", NAME)

	//-- assert
	s.RequireResultPass()
	assert.Contains(s.T(), out.ReadLine(), fmt.Sprintf("Hello %s", NAME))
}
