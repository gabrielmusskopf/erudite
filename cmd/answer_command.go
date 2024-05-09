package main

import (
	"flag"
	"fmt"
)

type AnswerCommandRequest struct {
	QuestionId int `json:"questionId"`
	AnswerId   int `json:"answerId"`
}

type AnswerCommand struct {
	flagSet *flag.FlagSet

	QuestionIdFlag *Flag[int]
	AnswerIdFlag   *Flag[int]
}

func NewAnswerCommand() AnswerCommand {
	cmd := AnswerCommand{}
	flagSet := flag.NewFlagSet(cmd.Name(), flag.ExitOnError)

	questionId := &Flag[int]{ShortName: "q", LongName: "question-id", Usage: "Question identifier."}
	answerId := &Flag[int]{ShortName: "a", LongName: "answer-id", Usage: "Answer identifier."}

	flagSet.IntVar(&questionId.Value, questionId.LongName, 0, questionId.Usage)
	flagSet.IntVar(&answerId.Value, answerId.LongName, 0, answerId.Usage)

	cmd.flagSet = flagSet
	cmd.QuestionIdFlag = questionId
	cmd.AnswerIdFlag = answerId

	flagSet.Usage = cmd.Usage

	return cmd
}

func (c AnswerCommand) Name() string {
	return "answer"
}

func (c AnswerCommand) Run(args []string) error {
	err := c.flagSet.Parse(args)
	if err != nil {
		return err
	}

	if c.QuestionIdFlag.Value == 0 {
		return fmt.Errorf("question identifier is missing")
	}

	if c.AnswerIdFlag.Value == 0 {
		return fmt.Errorf("answer identifier is missing")
	}

	request := AnswerCommandRequest{
		QuestionId: c.QuestionIdFlag.Value,
		AnswerId:   c.AnswerIdFlag.Value,
	}

	resp, err := server.Post("/answers", request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return nil
}

func (c AnswerCommand) Description() string {
	return "Answer question."
}

func (c AnswerCommand) Usage() {
	fmt.Printf(
		"%s\n\n%s <text>:\n%s\n%s\n",
		c.Description(),
		c.Name(),
		c.QuestionIdFlag.FullUsage(),
		c.AnswerIdFlag.FullUsage())
}
