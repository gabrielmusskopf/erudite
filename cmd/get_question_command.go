package main

import (
	"flag"
	"fmt"
	"io"
)

type GetQuestionCommand struct {
	flagSet *flag.FlagSet

	Id string
}

func NewGetQuestionCommand() Command {
	cmd := GetQuestionCommand{}
	cmd.flagSet = flag.NewFlagSet(cmd.Name(), flag.ExitOnError)
	cmd.flagSet.Usage = cmd.Usage
	return cmd
}

func (c GetQuestionCommand) Name() string {
	return "get-question"
}

func (c GetQuestionCommand) Run(args []string) error {
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

	//TODO: Change this panicking to something better
	resp, err := server.Get("/questions/" + c.Id)
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

func (c GetQuestionCommand) Description() string {
	return "Get question with provided identifier"
}

func (c GetQuestionCommand) Usage() {
	fmt.Printf(
		"%s\n\n%s <id>\n",
		c.Description(),
		c.Name())
}
