package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
)

type GetQuestionCommand struct {
	flagSet *flag.FlagSet

	Id      string
	TagFlag *Flag[Flags]
}

func NewGetQuestionCommand() Command {
	cmd := GetQuestionCommand{}
	flagSet := flag.NewFlagSet(cmd.Name(), flag.ExitOnError)

	tag := &Flag[Flags]{ShortName: "t", LongName: "tag", Usage: "Question related tags. Optional."}

	flagSet.Var(&tag.Value, tag.LongName, tag.Usage)
	flagSet.Var(&tag.Value, tag.ShortName, tag.Usage)

	cmd.flagSet = flagSet
	cmd.TagFlag = tag

	flagSet.Usage = cmd.Usage

	return cmd
}

func (c GetQuestionCommand) Name() string {
	return "get-question"
}

func (c GetQuestionCommand) Run(args []string) error {
	c.flagSet.Parse(args)
	c.Id = c.flagSet.Arg(0)

	var resp *http.Response

	if c.Id != "" {
		resp = getQuestionById(c.Id)
	} else if len(c.TagFlag.Value) != 0 {
		resp = getQuestionByTags(c.TagFlag.Value)
	} else {
		return fmt.Errorf("id or additional parameters are missing")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println(string(body))

	return nil
}

func getQuestionById(id string) *http.Response {
	resp, err := server.Get("/questions/" + id)
	if err != nil {
		//TODO: Change this panicking to something better
		panic(err)
	}
	return resp
}

func getQuestionByTags(tags []string) *http.Response {
	query := "?"
	for i, tag := range tags {
		query += "tag=" + tag
		if i != len(tags)-1 {
			query += "&"
		}
	}
	resp, err := server.Get("/questions/by" + query)
	if err != nil {
		//TODO: Change this panicking to something better
		panic(err)
	}
	return resp
}

func (c GetQuestionCommand) Description() string {
	return "Get question with provided identifier or tags. Important, question identifier have precedence over other tags."
}

func (c GetQuestionCommand) Usage() {
	fmt.Printf(
		"%s\n\n%s <id>:\n%s\n",
		c.Description(),
		c.Name(),
		c.TagFlag.FullUsage(),
	)
}
