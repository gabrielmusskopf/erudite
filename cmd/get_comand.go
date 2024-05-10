package main

import (
	"flag"
	"fmt"
)

type GetSubCommand struct {
	flagSet *flag.FlagSet

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
	flagSet.Usage = func() {
		fmt.Println(cmd.Description() + "\n")
		cmd.Usage()
	}

	cmd.flagSet = flagSet

	return cmd
}

func (c GetSubCommand) Name() string {
	return "get"
}

func (c GetSubCommand) Run(args []string) error {
	c.flagSet.Parse(args)

	if len(c.flagSet.Args()) < 1 {
		return userErr{msg: "one subcommand must be informed", command: c}
	}

	if err := CheckServer(c.flagSet.Args()); err != nil {
		return err
	}

	executedCmd, err := FindAndRunCommand(c.flagSet.Args(), c.commands)
	if err != nil {
		if executedCmd == nil {
			return parseUserErr(c, err)
		}
		return parseUserErr(executedCmd, err)
	}

	return nil
}

func (c GetSubCommand) Description() string {
	return "Get resources from Erudite server"
}

func (c GetSubCommand) Commands() []Command {
	return c.commands
}

func (c GetSubCommand) Usage() {
	CommandsUsage(c)
}
