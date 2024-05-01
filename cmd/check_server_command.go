package main

import (
	"flag"
	"fmt"
	"net/http"
)

type CheckServerCommand struct {
	flagSet *flag.FlagSet
}

func NewCheckServerCommand() Command {
	cmd := CheckServerCommand{}
	cmd.flagSet = flag.NewFlagSet(cmd.Name(), flag.ExitOnError)
	cmd.flagSet.Usage = cmd.Usage
	return cmd
}

func (c CheckServerCommand) Name() string {
	return "check-server"
}

func (c CheckServerCommand) Run(args []string) error {
	err := c.flagSet.Parse(args)
	if err != nil {
		return err
	}

	_, err = http.Get(server.URL + "/ping")
	if err != nil {
		return fmt.Errorf("Server is unreachable: %s", server.URL)
	}

	return nil
}

func (c CheckServerCommand) Description() string {
	return "Check if server is running and responding"
}

func (c CheckServerCommand) Usage() {
	fmt.Printf(
		"%s\n\n %s\n",
		c.Description(),
		c.Name())
}
