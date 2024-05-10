package main

import (
	"fmt"
	"os"
	"strings"
)

type userErr struct {
	msg     string
	command Command
}

func (e userErr) Error() string {
	return e.msg
}

func parseUserErr(cmd Command, err error) error {
	return userErr{msg: err.Error(), command: cmd}
}

var server = Server{URL: "http://localhost:8080"}
var checkServer = NewCheckServerCommand()

// Single end command
type Command interface {
	Name() string
	Description() string
	Run([]string) error
	Usage()
}

// Command aggregator. Like Command, but with n subcommands
type SubCommandable interface {
	Command
	Commands() []Command
}

func CommandsUsage(command SubCommandable) {
	fmt.Printf("Usage: %s [command] [flags|arguments]\n\n", command.Name())
	fmt.Print("Commands:\n")
	for _, cmd := range command.Commands() {
		space := 15 - len(cmd.Name())
		fmt.Printf("  %s %s %s\n", cmd.Name(), strings.Repeat(" ", space), cmd.Description())
	}
	fmt.Printf("\nTip: %s [command] -h\n", command.Name())
}

func containsHelp(args []string) bool {
	for _, arg := range args {
		if arg == "-h" || arg == "--help" {
			return true
		}
	}
	return false
}

func CheckServer(args []string) error {
	if !containsHelp(args) {
		// only command help will be executed
		if err := checkServer.Check(); err != nil {
			return err
		}
	}
	return nil
}

func FindAndRunCommand(args []string, subcommands []Command) (Command, error) {
	name := args[0]
	for _, cmd := range subcommands {
		if name == cmd.Name() {
			return cmd, cmd.Run(args[1:])
		}
	}

	return nil, fmt.Errorf("unknown command: %s", name)
}

func main() {
	eructl := NewEructlCommand()

	if err := eructl.Run(os.Args[1:]); err != nil {
		fmt.Printf("ERROR: %v\n\n", err.Error())

		switch e := err.(type) {
		case userErr:
			e.command.Usage()
		}
	}
}
