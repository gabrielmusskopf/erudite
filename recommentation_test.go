package main

import (
	"slices"
	"testing"
)

func TestBasedByWorstTagRecommendation(t *testing.T) {
	wantTag := "MATH"
	wantLimit := 1
	QuestionDB = NoopQuestionDatabase{}
	AnswerDB = NoopAnswerDatabase{}

	got := BasedOnWorstTag(RecommendationOptions{limit: wantLimit})

	if len(got) != wantLimit {
		t.Errorf("got %v questions but want 1", len(got))
	}

	question := got[0]
	if !slices.Contains(question.Tags, wantTag) {
		t.Errorf("got tag %v but want %s", question.Id, wantTag)
	}
}

func TestBasedByWorstTagRecommendationLimit2(t *testing.T) {
	wantTag := "MATH"
	wantLimit := 2
	QuestionDB = NoopQuestionDatabase{}
	AnswerDB = NoopAnswerDatabase{}

	got := BasedOnWorstTag(RecommendationOptions{limit: wantLimit})

	if len(got) != wantLimit {
		t.Errorf("got %v questions but want 1", len(got))
	}

	for _, question := range got {
		if !slices.Contains(question.Tags, wantTag) {
			t.Errorf("got tag %v but want %s", question.Id, wantTag)
		}
	}
}

type NoopQuestionDatabase struct {
}

func (d NoopQuestionDatabase) GetAny(GetQuestionOptions) ([]Question, error) {
	return []Question{
		{Id: 1, Text: "Question 1", Tags: []string{"MATH"}},
		{Id: 2, Text: "Question 2", Tags: []string{"MATH"}},
		{Id: 3, Text: "Question 3", Tags: []string{"MATH"}},
		{Id: 4, Text: "Question 4", Tags: []string{"GO"}},
		{Id: 5, Text: "Question 5", Tags: []string{"GO"}},
	}, nil
}

func (d NoopQuestionDatabase) Save(*Question) {
}

func (d NoopQuestionDatabase) Get(int) (*Question, error) {
	return nil, nil
}

type NoopAnswerDatabase struct {
}

func (d NoopAnswerDatabase) GetQuestionAnswered(questionId int) ([]Answer, error) {
	answers := map[int][]Answer{
		1: {{Id: 1, Text: "Answer 1", Correct: false}},
		2: {{Id: 2, Text: "Answer 1", Correct: false}},
		3: {{Id: 3, Text: "Answer 3", Correct: true}},
		4: {{Id: 4, Text: "Answer 4", Correct: false}},
		5: {{Id: 5, Text: "Answer 5", Correct: true}},
	}

	return answers[questionId], nil
}
func (d NoopAnswerDatabase) RegisterAnswer(questionId, answerId int) error {
	return nil
}
