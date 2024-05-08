package main

import (
	"database/sql"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

type QuestionDatabase interface {
	Save(*Question)
	Get(int) *Question
}

type questionDB struct {
	db *sql.DB
}

func (p questionDB) Save(question *Question) {
	tagIds := p.insertTags(question.Tags)

	var questionId int
	err := p.db.
		QueryRow("INSERT INTO questions (text, deleted, creation_date) VALUES ($1, $2, $3) RETURNING id",
			question.Text, false, time.Now()).
		Scan(&questionId)

	if err != nil {
		panic(err)
	}

	for _, tagId := range tagIds {
		var id int
		err = p.db.
			QueryRow("INSERT INTO question_tags (question_id, tag_id) VALUES ($1, $2) RETURNING id",
				questionId, tagId).
			Scan(&id)

		if err != nil {
			panic(err)
		}
	}

	for _, answer := range question.Answers {
		var id int
		err = p.db.
			QueryRow("INSERT INTO answers (question_id, text, correct) VALUES ($1, $2, $3) RETURNING id",
				questionId, answer.Text, answer.Correct).
			Scan(&id)

		if err != nil {
			panic(err)
		}
	}

	question.Id = questionId
}

func (p *questionDB) insertTags(tags []string) []int {
	tagIds := make([]int, 0)

	for _, tag := range tags {
		tag = strings.ToUpper(tag)
		var tagId int

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

func (p questionDB) Get(id int) *Question {
	var questionText string
	var questionCreationDate time.Time
	err := p.db.
		QueryRow("select text, creation_date from questions where id = $1 and deleted = false", id).
		Scan(&questionText, &questionCreationDate)

	if err != nil {
		panic(err)
	}

	return &Question{
		Id:           id,
		Text:         questionText,
		CreationDate: questionCreationDate,
		Tags:         p.getTags(id),
		Answers:      p.getAnswers(id),
	}
}

func (p questionDB) getTags(id int) []string {
	rows, err := p.db.Query(
		`select text from tags t 
        join question_tags q on t.id = q.tag_id
        where q.question_id = $1`, id)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	tags := make([]string, 0)
	for rows.Next() {
		var text string
		err = rows.Scan(&text)
		if err != nil {
			panic(err)
		}
		tags = append(tags, text)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return tags
}

func (p questionDB) getAnswers(id int) []Answer {
	rows, err := p.db.
		Query(`select id, text, correct from answers where question_id = $1`, id)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	answers := make([]Answer, 0)
	for rows.Next() {
		var id int
		var text string
		var correct bool
		err = rows.Scan(&id, &text, &correct)
		if err != nil {
			panic(err)
		}
		answers = append(answers, Answer{Id: id, Text: text, Correct: correct})
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return answers
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
