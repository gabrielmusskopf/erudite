package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type QuestionHandler struct {
}

type ResponseID struct {
	Id int `json:"id"`
}

func (h *QuestionHandler) HandleQuestionCreation(w http.ResponseWriter, r *http.Request) {
	question := &Question{}
	if err := json.NewDecoder(r.Body).Decode(&question); err != nil {
		writeError("could not decode request body", 400, w)
		return
	}
	fmt.Printf("New question received (%+v)\n", question)

	if len(question.Text) == 0 {
		writeError("question should have a text", 400, w)
		return
	}

	if len(question.Tags) == 0 {
		writeError("question should have at least one tag associated", 400, w)
		return
	}

	if len(question.Answers) == 0 {
		writeError("question should have at least one answer associated", 400, w)
		return
	}

	rigth := 0
	for _, answer := range question.Answers {
		if answer.Correct {
			rigth++
		}
		if rigth > 1 {
			writeError("must have exactly one rigth answer", 400, w)
			return
		}
	}
	if rigth == 0 {
		writeError("must have one rigth answer", 400, w)
		return
	}

	QuestionDB.Save(question)
	write(ResponseID{Id: question.Id}, w)
	fmt.Println("New question saved")
}

func (h *QuestionHandler) HandleQuestionGet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeError("id is not present", 400, w)
		return
	}

	query := r.URL.Query()
	if !query.Has("tag") {
	}

	question, err := QuestionDB.Get(id)
	if err != nil {
		writeError(err.Error(), 404, w)
		return
	}

	write(question, w)
}

func (h *QuestionHandler) HandleQuestionGetAny(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	question, err := QuestionDB.GetAny(GetQuestionOptions{
		tags: query["tag"],
	})
	if err != nil {
		writeError(err.Error(), 400, w)
		return
	}

	write(question, w)
}
