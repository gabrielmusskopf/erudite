package main

import (
	"flag"
	"fmt"
	"io"
)

type GetAnswersCommand struct {
	flagSet *flag.FlagSet

	Id string
}

func NewGetAnswersCommand() Command {
	cmd := GetAnswersCommand{}
	cmd.flagSet = flag.NewFlagSet(cmd.Name(), flag.ExitOnError)
	cmd.flagSet.Usage = cmd.Usage
	return cmd
}

func (c GetAnswersCommand) Name() string {
	//TODO: Implement get with subcommands.
	//Like "eructl get answer ..." and "eructl get question ..."
	return "get-answers"
}

func (c GetAnswersCommand) Run(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("id is missing")
	}

	if args[0] == "-h" || args[0] == "--help" {
		return c.flagSet.Parse(args)
	}

	c.Id = args[0]

	// search for help
	err := c.flagSet.Parse(args[1:])
	if err != nil {
		return err
	}

	//TODO: Change this panicking to something better
	resp, err := server.Get("/answers/question/" + c.Id)
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println(string(body))

	return nil
}

func (c GetAnswersCommand) Description() string {
	return "Get question answers"
}

func (c GetAnswersCommand) Usage() {
	fmt.Printf(
		"%s\n\n%s <id>\n",
		c.Description(),
		c.Name())
}
