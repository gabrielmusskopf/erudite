package main

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

const ()

var QuestionDB Database[Question]

type Database[T any] interface {
	Save(*T)
	Get(int) *T
}

func configureDb() {
	QuestionDB = PostgreStart()
}

type PostgreSQL struct {
	db       *sql.DB
	host     string
	port     int
	user     string
	password string
	dbname   string
}

func PostgreStart() Database[Question] {
	pdb := &PostgreSQL{
		host:     "localhost",
		port:     5432,
		user:     "postgres",
		password: "postgres",
		dbname:   "postgres",
	}

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		pdb.host, pdb.port, pdb.user, pdb.password, pdb.dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	pdb.db = db

	return pdb
}

func (p *PostgreSQL) Save(question *Question) {
	tagIds := p.insertTags(question.Tags)

	var questionId int
	err := p.db.
		QueryRow("INSERT INTO questions (text, deleted, creationdate) VALUES ($1, $2, $3) RETURNING id",
			question.Text, false, time.Now()).
		Scan(&questionId)

	if err != nil {
		panic(err)
	}

	for _, tagId := range tagIds {
		var id int
		err = p.db.
			QueryRow("INSERT INTO questionstags (questionid, tagid) VALUES ($1, $2) RETURNING id",
				questionId, tagId).
			Scan(&id)

		if err != nil {
			panic(err)
		}
	}

	for _, answer := range question.Answers {
		var id int
		err = p.db.
			QueryRow("INSERT INTO answers (questionid, text, correct) VALUES ($1, $2, $3) RETURNING id",
				questionId, answer.Text, answer.Correct).
			Scan(&id)

		if err != nil {
			panic(err)
		}
	}

	question.Id = questionId
}

func (p *PostgreSQL) insertTags(tags []string) []int {
	tagIds := make([]int, 0)

	for _, tag := range tags {
		var tagId int
		tag = strings.ToUpper(tag)

		err := p.db.
			QueryRow("SELECT id FROM tags WHERE text = $1", tag).
			Scan(&tagId)

		if err != nil {
			// it does not exist yet
			err := p.db.
				QueryRow("INSERT INTO tags (text) VALUES ($1) RETURNING id", tag).
				Scan(&tagId)

			if err != nil {
				panic(err)
			}
		}
		tagIds = append(tagIds, tagId)
	}

	return tagIds
}

func (p *PostgreSQL) Get(id int) *Question {
	fmt.Println(id)

	var text string
	err := p.db.
		QueryRow("select text from questions where id = $1", id).
		Scan(&text)

	if err != nil {
		panic(err)
	}

	return &Question{
		Id:   id,
		Text: text,
	}
}

type InMemoryDb struct {
	questions []*Question
}

func (db *InMemoryDb) Save(question *Question) {
	question.Id = len(db.questions) + 1
	question.Deleted = false
	question.CreationDate = time.Now()
	db.questions = append(db.questions, question)
}

func (db *InMemoryDb) Get(id int) *Question {
	for _, question := range db.questions {
		if question.Id == id {
			return question
		}
	}
	return nil
}
