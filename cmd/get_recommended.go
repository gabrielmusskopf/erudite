package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
)

type GetRecommendedQuestionCommand struct {
	flagSet *flag.FlagSet

	Limit      *Flag[int]
	ByWorstTag *Flag[bool]
}

func NewGetRecommendedCommand() Command {
	cmd := GetRecommendedQuestionCommand{}
	flagSet := flag.NewFlagSet(cmd.Name(), flag.ExitOnError)

	limit := &Flag[int]{ShortName: "l", LongName: "limit", Usage: "Limit number of recommended questions."}
	byWorstTag := &Flag[bool]{ShortName: "by-worst-tag", LongName: "", Usage: "Recommended questions based on worst tag score."}

	flagSet.IntVar(&limit.Value, limit.ShortName, 1, limit.Usage)
	flagSet.IntVar(&limit.Value, limit.LongName, 1, limit.Usage)
	flagSet.BoolVar(&byWorstTag.Value, byWorstTag.ShortName, false, byWorstTag.Usage)

	cmd.flagSet = flagSet
	cmd.ByWorstTag = byWorstTag
	cmd.Limit = limit

	flagSet.Usage = cmd.Usage

	return cmd
}

func (c GetRecommendedQuestionCommand) Name() string {
	return "recommend"
}

func (c GetRecommendedQuestionCommand) Run(args []string) error {
	c.flagSet.Parse(args)

	var resp *http.Response

	if c.ByWorstTag.Value {
		r, err := c.getByWorstTag()
		if err != nil {
			return err
		}
		resp = r
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println(string(body))

	return nil
}

func (c GetRecommendedQuestionCommand) getByWorstTag() (*http.Response, error) {
	url := fmt.Sprintf("/questions/by/worst/tag?limit=%d", c.Limit.Value)
	resp, err := server.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c GetRecommendedQuestionCommand) Description() string {
	return "Get recommended question with provided filters and options."
}

func (c GetRecommendedQuestionCommand) Usage() {
	fmt.Printf(
		"%s\n\n%s:\n%s\n%s\n",
		c.Description(),
		c.Name(),
		c.Limit.FullUsage(),
		c.ByWorstTag.FullUsage(),
	)
}
