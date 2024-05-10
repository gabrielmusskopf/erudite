package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

var server = Server{URL: "http://localhost:8080"}
var checkServer = NewCheckServerCommand()

type Command interface {
	Name() string
	Description() string
	Run([]string) error
}

type EructlCommand struct {
	flagSet  *flag.FlagSet
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
	flagSet.Usage = cmd.Usage

	cmd.flagSet = flagSet

	return cmd
}

func (c EructlCommand) Name() string {
	return "eructl"
}

func (c EructlCommand) Run(args []string) error {
	c.flagSet.Parse(args)

	if len(c.flagSet.Args()) < 1 {
		fmt.Print("ERROR: one command must be informed\n\n")
		c.Usage()
		return nil
	}

	if err := CheckServer(c.flagSet.Args()[1:]); err != nil {
		return err
	}

	if err := IterateCommands(c.flagSet.Args(), c.commands); err != nil {
		fmt.Println(err.Error())
		c.Usage()
	}

	return nil
}

func (c EructlCommand) Description() string {
	return ""
}

func (c EructlCommand) Usage() {
	CommandsUsage(c.Name(), c.commands)
}

func CommandsUsage(commandName string, commands []Command) {
	fmt.Printf("Usage: %s [command] [flags|arguments]\n\n", commandName)
	fmt.Print("Commands:\n")
	for _, cmd := range commands {
		space := 15 - len(cmd.Name())
		fmt.Printf("  %s %s %s\n", cmd.Name(), strings.Repeat(" ", space), cmd.Description())
	}
	fmt.Printf("\nTip: %s [command] -h\n", commandName)
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

func IterateCommands(args []string, subcommands []Command) error {
	name := args[0]
	for _, cmd := range subcommands {
		if name == cmd.Name() {
			args := args[1:]

			if err := cmd.Run(args); err != nil {
				fmt.Printf("ERROR: %v\n", err.Error())
			}
			return nil
		}
	}

	return fmt.Errorf("ERROR: unknown command: %s\n", name)
}

type Server struct {
	URL string
}

func (s Server) Post(path string, body any) (*http.Response, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", s.URL+path, bytes.NewReader(b))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{Timeout: 10 * time.Second}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	return res, nil
}

func (s Server) Get(path string) (*http.Response, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	r, err := http.Get(s.URL + path)
	if err != nil {
		return nil, err
	}

	return r, nil
}

type Flags []string

func (i *Flags) String() string {
	return fmt.Sprintf("%v", *i)
}

func (i *Flags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

type Flag[T any] struct {
	Usage     string
	ShortName string
	LongName  string
	Value     T
}

func (f Flag[T]) Names() (short, long string) {
	return f.ShortName, f.LongName
}

func (f Flag[T]) FullUsage() string {
	s := "  "
	short, long := f.Names()
	if short != "" {
		s += "-" + short
	}
	if long != "" {
		s += ", --" + long
	}
	s += ":" + strings.Repeat(" ", 25-len(s)) + f.Usage
	return s
}

func main() {
	eructl := NewEructlCommand()
	if err := eructl.Run(os.Args[1:]); err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
	}
}
