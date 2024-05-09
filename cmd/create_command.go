package main

import (
	"flag"
	"fmt"
	"io"
)

type CreateCommandRequest struct {
	Text    string   `json:"text"`
	Tags    []string `json:"tags"`
	Answers []Answer `json:"answers"`
}

type Answer struct {
	Text    string `json:"text"`
	Correct bool   `json:"correct"`
}

type CreateCommand struct {
	flagSet *flag.FlagSet

	Text             string
	TagFlag          *Flag[Flags]
	WrongAnswersFlag *Flag[Flags]
	RigthAnswerFlag  *Flag[string]
}

func NewCreateCommand() CreateCommand {
	cmd := CreateCommand{}
	flagSet := flag.NewFlagSet(cmd.Name(), flag.ExitOnError)

	tag := &Flag[Flags]{ShortName: "t", LongName: "tag", Usage: "Question related tags. Optional."}
	wrongAnswers := &Flag[Flags]{ShortName: "wa", LongName: "wrong-answer", Usage: "Question wrong answer(s)."}
	rigthAnswer := &Flag[string]{ShortName: "ra", LongName: "rigth-answer", Usage: "Question rigth answer."}

	flagSet.Var(&tag.Value, tag.LongName, tag.Usage)
	flagSet.Var(&tag.Value, tag.ShortName, tag.Usage)

	flagSet.Var(&wrongAnswers.Value, wrongAnswers.LongName, wrongAnswers.Usage)
	flagSet.Var(&wrongAnswers.Value, wrongAnswers.ShortName, wrongAnswers.Usage)

	flagSet.StringVar(&rigthAnswer.Value, rigthAnswer.LongName, "", rigthAnswer.Usage)
	flagSet.StringVar(&rigthAnswer.Value, rigthAnswer.ShortName, "", rigthAnswer.Usage)

	cmd.flagSet = flagSet
	cmd.TagFlag = tag
	cmd.WrongAnswersFlag = wrongAnswers
	cmd.RigthAnswerFlag = rigthAnswer

	flagSet.Usage = cmd.Usage

	return cmd
}

func (c CreateCommand) Name() string {
	return "create"
}

func (c CreateCommand) Run(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("text is missing")
	}

	if containsHelp(args) {
		return c.flagSet.Parse(args)
	}

	c.Text = args[0]

	err := c.flagSet.Parse(args[1:])
	if err != nil {
		return err
	}

	if len(c.WrongAnswersFlag.Value) == 0 {
		return fmt.Errorf("must have at least one wrong answer")
	}

	if len(c.RigthAnswerFlag.Value) == 0 {
		return fmt.Errorf("must have one rigth answer")
	}

	answers := make([]Answer, 0)
	answers = append(answers, Answer{Text: c.RigthAnswerFlag.Value, Correct: true})
	for _, wa := range c.WrongAnswersFlag.Value {
		answers = append(answers, Answer{Text: wa, Correct: false})
	}

	request := CreateCommandRequest{
		Answers: answers,
		Text:    c.Text,
		Tags:    c.TagFlag.Value,
	}

	resp, err := server.Post("/questions", request)
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

func (c CreateCommand) Description() string {
	return "Create questions with provided tags and answers."
}

func (c CreateCommand) Usage() {
	fmt.Printf(
		"%s\n\n%s <text>:\n%s\n%s\n%s\n",
		c.Description(),
		c.Name(),
		c.TagFlag.FullUsage(),
		c.WrongAnswersFlag.FullUsage(),
		c.RigthAnswerFlag.FullUsage())
}
