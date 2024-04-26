package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func checkServer() {
	_, err := http.Get("http://localhost:8080/ping")
	if err != nil {
		panic(err)
	}
}

var commands = []Command{
	NewCreateCommand(),
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
	s := ""
	short, long := f.Names()
	if short != "" {
		s += "-" + short
	}
	if long != "" {
		s += ", --" + long
	}
	s += ":\n    " + f.Usage + "\n"
	return s
}

type Command interface {
	Name() string
	Run([]string) error
	Usage() string
}

func usage() {
	fmt.Print("Usage: eructl [command] [flags|arguments]\n\n")
	for _, cmd := range commands {
		fmt.Print(cmd.Usage())
	}
}

func main() {
	flag.Usage = usage
	flag.Parse()

	checkServer()

	if len(os.Args) < 2 {
		fmt.Print("ERROR: one command must be informed\n\n")
		usage()
		return
	}

	command := os.Args[1]
	for _, cmd := range commands {
		if command == cmd.Name() {
			err := cmd.Run(os.Args[2:])
			if err != nil {
				fmt.Printf("ERROR: %v\n", err.Error())
			}
			return
		}
	}

	fmt.Printf("ERROR: unknown command: %s\n\n", command)
	usage()
}
