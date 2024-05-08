package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type AnswerHandler struct {
}

type QuestionAnswer struct {
	QuestionId int `json:"questionId"`
	AnswerId   int `json:"answerId"`
}

func (h *AnswerHandler) HandleQuestionAnswer(w http.ResponseWriter, r *http.Request) {
	questionAnswer := &QuestionAnswer{}
	if err := json.NewDecoder(r.Body).Decode(&questionAnswer); err != nil {
		writeError("could not decode request body", 400, w)
		return
	}
	fmt.Printf("New question answer received (%+v)\n", questionAnswer)

	question, err := QuestionDB.Get(questionAnswer.QuestionId)
	if err != nil {
		writeFullError("invalid question id", err.Error(), 400, w)
		return
	}

	if !containsAnswer(questionAnswer.AnswerId, *question) {
		writeError("invalid answer id", 400, w)
		return
	}

	if err := AnswerDB.RegisterAnswer(questionAnswer.QuestionId, questionAnswer.AnswerId); err != nil {
		writeError(err.Error(), 400, w)
	}
}

func containsAnswer(answerId int, question Question) bool {
	for _, a := range question.Answers {
		if a.Id == answerId {
			return true
		}
	}
	return false
}

func (h *AnswerHandler) HandleGetQuestionAnswers(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeError("question id is not present", 400, w)
		return
	}

	answers, err := AnswerDB.GetQuestionAnswered(id)
	if err != nil {
		log.Printf("ERROR: %v\n", err)
	}

	write(answers, w)
}
