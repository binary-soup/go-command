package command

import (
	"flag"
)

// A base type for commands.
type CommandBase struct {
	Name        string
	Description string
	Flags       *flag.FlagSet
}

// Create a new CommandBase from a command name and description.
func NewCommandBase(name, description string) CommandBase {
	return CommandBase{
		Name:        name,
		Description: description,
	}
}

// Return the command's id (ie. its name).
func (cmd CommandBase) GetID() string {
	return cmd.Name
}

// Return the command's usage string (ie. its description).
func (cmd CommandBase) GetUsage() string {
	return cmd.Description
}

// Does nothing.
func (cmd CommandBase) Initialize() error {
	return nil
}
