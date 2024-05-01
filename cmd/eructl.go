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
var commands = []Command{
	NewCreateCommand(),
	NewGetCommand(),
	checkServer,
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

type Command interface {
	Name() string
	Description() string
	Run([]string) error
}

func usage() {
	fmt.Print("Usage: eructl [command] [flags|arguments]\n\n")
	fmt.Print("Commands:\n")
	for _, cmd := range commands {
		space := 15 - len(cmd.Name())
		fmt.Printf("  %s %s %s\n", cmd.Name(), strings.Repeat(" ", space), cmd.Description())
	}
	fmt.Print("\nTip: eructl [command] -h\n")
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Print("ERROR: one command must be informed\n\n")
		usage()
		return
	}

	command := os.Args[1]
	for _, cmd := range commands {
		if command == cmd.Name() {
			args := os.Args[2:]

			if err := checkServer.Run(args); err != nil {
				fmt.Printf("ERROR: %v\n", err.Error())
				return
			}

			if cmd == checkServer {
				fmt.Println("pong")
				return
			}

			if err := cmd.Run(args); err != nil {
				fmt.Printf("ERROR: %v\n", err.Error())
			}
			return
		}
	}

	fmt.Printf("ERROR: unknown command: %s\n\n", command)
	usage()
}
