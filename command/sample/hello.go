package sample

import (
	"errors"
	"fmt"

	"github.com/binarysoupdev/go-commando/command"
)

// Sample hello command for printing "Hello" to the console.
type HelloCommand struct {
	command.CommandBase
	command.FlagCommand
}

func NewHelloCommand() *HelloCommand {
	return &HelloCommand{
		CommandBase: command.NewCommandBase("hello", "prints \"Hello {name}\" to the console"),
	}
}

func (cmd *HelloCommand) Initialize() error {
	cmd.InitFlagSet(cmd.Name, cmd.Description)
	return nil
}

func (cmd HelloCommand) Run(args []string) error {
	name := cmd.Flags.String("name", "World", "name to use when saying hello")
	cmd.ParseFlags(args)

	if *name == "" {
		return errors.New("name cannot be empty")
	}

	fmt.Printf("Hello %s!\n", *name)
	return nil
}
