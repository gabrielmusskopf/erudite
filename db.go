package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var QuestionDB QuestionDatabase
var AnswerDB AnswerDatabase

func configureDb() {
	db := PostgreStart()
	QuestionDB = questionDB{db: db}
	AnswerDB = answerDB{db: db}
}

type DBConnection struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
}

func PostgreStart() *sql.DB {
	dbc := &DBConnection{
		host:     "localhost",
		port:     5432,
		user:     "postgres",
		password: "postgres",
		dbname:   "postgres",
	}

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		dbc.host, dbc.port, dbc.user, dbc.password, dbc.dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
