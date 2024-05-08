package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var questionHandler = &QuestionHandler{}
var answerHandler = &AnswerHandler{}

type ResponseError struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
	Detail  string `json:"detail,omitempty"`
}

func configureHandlers(mux *http.ServeMux) {
	// questions
	mux.Handle("/questions", Post(questionHandler.HandleQuestionCreation))
	mux.Handle("/questions/{id}", Get(questionHandler.HandleQuestionGet))

	// answers
	mux.Handle("/answers", Post(answerHandler.HandleQuestionAnswer))

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Pong"))
	})
}

func writeError(msg string, status int, w http.ResponseWriter) {
	writeFullError(msg, "", status, w)
}

func writeFullError(msg, detail string, status int, w http.ResponseWriter) {
	error := &ResponseError{
		Status:  status,
		Detail:  detail,
		Message: msg,
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

type Post http.HandlerFunc
type Get http.HandlerFunc
type Put http.HandlerFunc
type Delete http.HandlerFunc

func (f Post) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		writeError("", http.StatusMethodNotAllowed, w)
		return
	}
	f(w, r)
}

func (f Get) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		writeError("", http.StatusMethodNotAllowed, w)
		return
	}
	f(w, r)
}

func (f Put) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		writeError("", http.StatusMethodNotAllowed, w)
		return
	}
	f(w, r)
}

func (f Delete) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		writeError("", http.StatusMethodNotAllowed, w)
		return
	}
	f(w, r)
}
