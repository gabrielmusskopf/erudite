package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type AnswerDatabase interface {
	Save(Answer)
	Get(int) *Answer
}

type answerDB struct {
	db *sql.DB
}

func (p answerDB) Save(answer Answer) {
}

func (p answerDB) Get(id int) *Answer {
	return nil
}
