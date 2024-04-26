package main

import "net/http"

var questionHandler = &QuestionHandler{}

func configureHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/questions/", questionHandler.HandleQuestionCreation)
	mux.HandleFunc("/questions/{id}", questionHandler.HandleQuestionGet)
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Pong"))
	})
}
