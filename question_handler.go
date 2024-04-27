package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type QuestionHandler struct {
}

type ResponseError struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
}

type ResponseID struct {
	Id int `json:"id"`
}

func writeError(msg string, status int, w http.ResponseWriter) {
	error := &ResponseError{
		Status: status,
	}
	if len(msg) != 0 {
		error.Message = msg
	}
	b, err := json.Marshal(error)
	if err != nil {
		panic(fmt.Sprintf("could not marshal %v\n", err))
	}
	http.Error(w, string(b), status)
}

func write[V any](v V, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		panic("could not serialize response")
	}
}

func (h *QuestionHandler) HandleQuestionCreation(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		writeError("", http.StatusMethodNotAllowed, w)
		return
	}

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

	question := QuestionDB.Get(id)
	if question == nil {
		writeError("", 404, w)
		return
	}

	write(question, w)
}
