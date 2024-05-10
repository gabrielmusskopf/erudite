package main

import (
	"flag"
	"fmt"
	"strings"
)

type GetSubCommand struct {
	flagSet  *flag.FlagSet
	commands []Command
}

func NewGetSubCommand() GetSubCommand {
	cmd := GetSubCommand{
		commands: []Command{
			NewGetQuestionCommand(),
			NewGetAnswersCommand(),
		},
	}

	flagSet := flag.NewFlagSet(cmd.Name(), flag.ExitOnError)
	flagSet.Usage = cmd.Usage

	cmd.flagSet = flagSet

	return cmd
}

func (c GetSubCommand) Name() string {
	return "get"
}

func (c GetSubCommand) Run(args []string) error {
	c.flagSet.Parse(args)

	if len(c.flagSet.Args()) < 2 {
		fmt.Print("ERROR: one subcommand must be informed\n\n")
		c.Usage()
		return nil
	}

	if !containsHelp(c.flagSet.Args()) {
		// only command help will be executed
		if err := checkServer.Check(); err != nil {
			fmt.Printf("ERROR: %v\n", err.Error())
			return nil
		}
	}

	if err := IterateCommands(c.flagSet.Args(), c.commands); err != nil {
		fmt.Println(err.Error())
		c.Usage()
	}

	return nil
}

func (c GetSubCommand) Description() string {
	return "Get subcommand"
}

func (c GetSubCommand) Usage() {
	fmt.Printf("Usage: eructl %s [command] [flags|arguments]\n\n", c.Name())
	fmt.Print("Commands:\n")
	for _, cmd := range c.commands {
		space := 15 - len(cmd.Name())
		fmt.Printf("  %s %s %s\n", cmd.Name(), strings.Repeat(" ", space), cmd.Description())
	}
}
