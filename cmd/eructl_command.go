package main

import (
	"flag"
	"fmt"
)

type EructlCommand struct {
	flagSet *flag.FlagSet

	commands []Command
}

func NewEructlCommand() EructlCommand {
	cmd := EructlCommand{
		commands: []Command{
			checkServer,
			NewCreateCommand(),
			NewAnswerCommand(),
			NewGetSubCommand(),
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

func (c EructlCommand) Name() string {
	return "eructl"
}

func (c EructlCommand) Run(args []string) error {
	c.flagSet.Parse(args)

	if len(c.flagSet.Args()) < 1 {
		return userErr{msg: "one subcommand must be informed\n", command: c}
	}

	if err := CheckServer(c.flagSet.Args()[1:]); err != nil {
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

func (c EructlCommand) Description() string {
	return "Interact with Erudite server"
}

func (c EructlCommand) Usage() {
	CommandsUsage(c)
}

func (c EructlCommand) Commands() []Command {
	return c.commands
}
