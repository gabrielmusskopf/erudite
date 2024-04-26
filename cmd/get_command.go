package main

import (
	"flag"
	"fmt"
)

type GetCommand struct {
	flagSet *flag.FlagSet

	Id string
}

func NewGetCommand() Command {
	cmd := GetCommand{}
	cmd.flagSet = flag.NewFlagSet(cmd.Name(), flag.ExitOnError)
	cmd.flagSet.Usage = cmd.Usage
	return cmd
}

func (c GetCommand) Name() string {
	return "get"
}

func (c GetCommand) Run(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("id is missing")
	}

	if args[0] == "-h" || args[0] == "--help" {
		return c.flagSet.Parse(args)
	}

	c.Id = args[0]

	err := c.flagSet.Parse(args[1:]) // search for help
	if err != nil {
		return err
	}

	fmt.Printf("id=%v\n", c.Id)
	return nil
}

func (c GetCommand) Description() string {
	return "Get question with provided identifier"
}

func (c GetCommand) Usage() {
	fmt.Printf(
		"%s\n\n%s <id>\n",
		c.Description(),
		c.Name())
}
