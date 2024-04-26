package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Question struct {
	Id           int       `json:"id"`
	Text         string    `json:"text"`
	Tags         []string  `json:"tags"`
	Deleted      bool      `json:"deleted"`
	CreationDate time.Time `json:"creationDate"`
	Answers      []Answer  `json:"answers"`
}

type Answer struct {
	Text    string `json:"text"`
	Correct bool   `json:"correct"`
}

type QuestionAnswered struct {
	Question Question
	Answer   Answer
}

func main() {
	mux := http.NewServeMux()
	configureHandlers(mux)
	configureDb()

	fmt.Println("Listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
