package main

import (
	"flag"
	"fmt"
)

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
	flagSet.Usage = func() { fmt.Print(cmd.Usage()) }

	tag := &Flag[Flags]{ShortName: "t", LongName: "tag", Usage: "question related tags"}
	wrongAnswers := &Flag[Flags]{ShortName: "wa", LongName: "wrong-answer", Usage: "question wrong answer(s)"}
	rigthAnswer := &Flag[string]{ShortName: "ra", LongName: "rigth-answer", Usage: "question rigth answer"}

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

	return cmd
}

func (c CreateCommand) Name() string {
	return "create"
}

func (c CreateCommand) Run(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("text is missing")
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

	fmt.Printf("text=%v\n", c.Text)
	fmt.Printf("tags=%v\n", c.TagFlag.Value)
	fmt.Printf("rigth-answer=%v\n", c.RigthAnswerFlag.Value)
	fmt.Printf("wrong-answer=%v\n", c.WrongAnswersFlag.Value)

	return nil
}

func (c CreateCommand) Usage() string {
	return fmt.Sprintf(
		"%s <text>:\n%s\n  %s\n  %s\n",
		c.Name(),
		c.TagFlag.FullUsage(),
		c.WrongAnswersFlag.FullUsage(),
		c.RigthAnswerFlag.FullUsage())
}
