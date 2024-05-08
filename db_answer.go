package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type AnswerDatabase interface {
	RegisterAnswer(questionId, answerId int)
}

type answerDB struct {
	db *sql.DB
}

func (p answerDB) RegisterAnswer(questionId, answerId int) {
	var id int
	err := p.db.
		QueryRow("INSERT INTO question_answers (question_id, answer_id) VALUES ($1, $2) RETURNING id", questionId, answerId).
		Scan(&id)

	if err != nil {
		panic(err)
	}
}
