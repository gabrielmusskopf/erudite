package main

import "net/http"

var questionHandler = &QuestionHandler{}

func configureHandlers(mux *http.ServeMux) {
	http.HandleFunc("/questions", questionHandler.HandleQuestionCreation)
	http.HandleFunc("/questions/{id}", questionHandler.HandleQuestionGet)
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Pong"))
	})
}
