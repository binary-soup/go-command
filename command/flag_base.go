package command

import (
	"flag"
	"fmt"

	"github.com/binarysoupdev/got-style/style"
)

// A base type for commands that use a flagset to parse arguments.
type FlagCommandBase struct {
	CommandBase
	Flags *flag.FlagSet
}

// Create a new FlagCommandBase from a command name and description.
func NewFlagCommandBase(name, description string) FlagCommandBase {
	return FlagCommandBase{
		CommandBase: CommandBase{
			Name:        name,
			Description: description,
		},
	}
}

// Create a new flagset and override its usage function.
// Never returns an error.
func (cmd *FlagCommandBase) Initialize() error {
	cmd.Flags = flag.NewFlagSet(cmd.Name, flag.ExitOnError)
	cmd.Flags.Usage = func() {
		fmt.Println(cmd.Description)
		style.New(style.MAGENTA).Println("Options:")
		cmd.Flags.PrintDefaults()
	}

	return nil
}

func (cmd FlagCommandBase) ParseFlags(args []string) {
	cmd.Flags.Parse(args)
}
