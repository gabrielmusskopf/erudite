package main

import (
	"flag"
	"fmt"
	"net/http"
)

func checkServer() {
	_, err := http.Get("http://localhost:8080/ping")
	if err != nil {
		//panic(err)
		//not panic for now...
		fmt.Println(err)
	}
}

/*
eruadm
erucli
eructl

text string
tags []string
answer Answer

Answer {
    text string
    correct bool
}

eructl "My question here"
    --tag "GO"
    --tag "Programming"
    --answer "Answer here"
    --rigth-answer "Answer here"
*/

type Flags []string

var (
	tags         Flags
	wrongAnswers Flags
	rigthAnswer  string
)

func (i *Flags) String() string {
	return fmt.Sprintf("%v", *i)
}

func (i *Flags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func main() {
	checkServer()

	flag.Var(&tags, "tag", "question related tags")
	flag.Var(&tags, "t", "question related tags")
	flag.Var(&wrongAnswers, "wrong-answer", "question wrong answer(s)")
	flag.Var(&wrongAnswers, "wa", "question wrong answer(s)")
	flag.StringVar(&rigthAnswer, "rigth-answer", "", "question rigth answer")
	flag.StringVar(&rigthAnswer, "ra", "", "question rigth answer")
	flag.Parse()

	fmt.Printf("%v\n", tags)
	fmt.Printf("%v\n", wrongAnswers)
	fmt.Printf("%v\n", rigthAnswer)
}
