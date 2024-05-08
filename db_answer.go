package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type AnswerDatabase interface {
	GetQuestionAnswered(questionId int) ([]Answer, error)
	RegisterAnswer(questionId, answerId int) error
}

type answerDB struct {
	db *sql.DB
}

func (p answerDB) RegisterAnswer(questionId, answerId int) error {
	var id int
	err := p.db.
		QueryRow("INSERT INTO question_answers (question_id, answer_id) VALUES ($1, $2) RETURNING id", questionId, answerId).
		Scan(&id)

	if err != nil {
		return err
	}

	return nil
}

func (a answerDB) GetQuestionAnswered(questionId int) ([]Answer, error) {
	answers := make([]Answer, 0)

	rows, err := a.db.Query(
		`select a.id, a."text", a.correct from answers a
        join question_answers qa ON qa.answer_id = a.id
        where qa.question_id = $1`, questionId)

	if err != nil {
		return answers, err
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var text string
		var correct bool
		err = rows.Scan(&id, &text, &correct)
		if err != nil {
			return answers, err
		}
		answers = append(answers, Answer{Id: id, Text: text, Correct: correct})
	}
	err = rows.Err()
	if err != nil {
		return answers, err
	}

	return answers, nil
}
