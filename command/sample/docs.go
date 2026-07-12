// Package sample provides sample commands that demonstrate the command package.
//
// Boiler plate command for convenience:
//
//	type SampleCommand struct {
//		command.CommandBase
//	}
//
//	func NewSampleCommand() *SampleCommand {
//		return &SampleCommand{
//			CommandBase: command.NewCommandBase("name", "description"),
//		}
//	}
//
//	func (cmd SampleCommand) Run(args []string) error {
//		return nil
//	}
package sample
