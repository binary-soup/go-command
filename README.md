# go-commando [![GoDoc](https://godoc.org/github.com/binarysoupdev/go-commando?status.svg)](https://pkg.go.dev/github.com/binarysoupdev/go-commando)

The `go-commando` module provides tools and testing utilities for creating and testing modular types (aka. commands) that encapsulate distinct program flow. Inspired by the _separation of concerns_ design principal.

## Basic Usage

### The Command Interface

```go
type Command interface {
	// Get the command's identifier.
	GetID() string

	// Get the command's usage string.
	GetUsage() string

	// Initializes the command before it's run and return any errors.
	Initialize() error

	// Run the command using the provided arguments and return any errors.
	Run(args []string) error
}
```

### A Sample Flag-Based Command

```go
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
```

### Sample `main` Function

```go
func main() {
	ls := flag.Bool("ls", false, "list all commands")
	flag.Parse()

	runner := command.NewRunner(
		sample.NewHelloCommand(),
	)

	if *ls || len(os.Args) < 2 {
		runner.ListCommands()
		return
	}

	if err := runner.RunCommand(os.Args[1], os.Args[2:]); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
```
