package main

import (
	"flag"
	"fmt"
	"net/http"
)

type CheckServerCommand struct {
	flagSet *flag.FlagSet
}

func NewCheckServerCommand() CheckServerCommand {
	cmd := CheckServerCommand{}
	cmd.flagSet = flag.NewFlagSet(cmd.Name(), flag.ExitOnError)
	cmd.flagSet.Usage = cmd.Usage
	return cmd
}

func (c CheckServerCommand) Name() string {
	return "check-server"
}

func (c CheckServerCommand) Check() error {
	_, err := http.Get(server.URL + "/ping")
	if err != nil {
		return fmt.Errorf("Server is unreachable: %s", server.URL)
	}
	return nil
}

func (c CheckServerCommand) Run(args []string) error {
	c.flagSet.Parse(args)
	if err := c.Check(); err != nil {
		return err
	}
	fmt.Println("Server is fine :)")
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
