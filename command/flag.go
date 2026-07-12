package command

import (
	"flag"
	"fmt"

	"github.com/binarysoupdev/got-style/style"
)

// Command component for parsing arguments using the flag package.
type FlagCommand struct {
	Flags *flag.FlagSet
}

// Create a new flagset and override its usage function.
func (cmd *FlagCommand) InitFlagSet(name, usage string) {
	cmd.Flags = flag.NewFlagSet(name, flag.ExitOnError)
	cmd.Flags.Usage = func() {
		fmt.Println(usage)
		style.New(style.MAGENTA).Println("Options:")
		cmd.Flags.PrintDefaults()
	}
}

// Shorthand for cmd.Flags.Parse(args).
func (cmd FlagCommand) ParseFlags(args []string) {
	cmd.Flags.Parse(args)
}
