package main

import (
	"flag"
	"fmt"
	"os"
)

type CreateCommand struct {
	flagSet          *flag.FlagSet
	TagFlag          Flag[Flags]
	WrongAnswersFlag Flag[Flags]
	RigthAnswerFlag  Flag[string]
}

func NewCreateCommand() CreateCommand {
	tag := Flag[Flags]{ShortName: "t", LongName: "tag", Usage: "question related tags"}
	wrongAnswers := Flag[Flags]{ShortName: "wa", LongName: "wrong-answer", Usage: "question wrong answer(s)"}
	rigthAnswer := Flag[string]{ShortName: "ra", LongName: "rigth-answer", Usage: "question rigth answer"}

	flagSet := flag.NewFlagSet("create", flag.ExitOnError)
	flagSet.Var(&tag.Value, tag.LongName, tag.Usage)
	flagSet.Var(&tag.Value, tag.ShortName, tag.Usage)

	flagSet.Var(&wrongAnswers.Value, wrongAnswers.LongName, wrongAnswers.Usage)
	flagSet.Var(&wrongAnswers.Value, wrongAnswers.ShortName, wrongAnswers.Usage)

	flagSet.StringVar(&rigthAnswer.Value, rigthAnswer.LongName, "", rigthAnswer.Usage)
	flagSet.StringVar(&rigthAnswer.Value, rigthAnswer.ShortName, "", rigthAnswer.Usage)

	cmd := CreateCommand{
		flagSet:          flagSet,
		TagFlag:          tag,
		WrongAnswersFlag: wrongAnswers,
		RigthAnswerFlag:  rigthAnswer,
	}

	flagSet.Usage = func() {
		fmt.Print(cmd.Usage())
	}

	return cmd
}

func (c CreateCommand) Name() string {
	return "create"
}

func (c CreateCommand) Run([]string) {
	err := c.flagSet.Parse(os.Args[2:])
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", c.TagFlag.Value)
	fmt.Printf("%v\n", c.RigthAnswerFlag.Value)
	fmt.Printf("%v\n", c.WrongAnswersFlag.Value)
}

func (c CreateCommand) Usage() string {
	return fmt.Sprintf(
		"%s:\n%s\n  %s\n  %s\n",
		c.Name(),
		c.TagFlag.FullUsage(),
		c.WrongAnswersFlag.FullUsage(),
		c.RigthAnswerFlag.FullUsage())
}
