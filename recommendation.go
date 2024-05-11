package main

import (
	"math/rand"
)

type RecommendationOptions struct {
	limit int
}

func groupByTag(questions []Question) map[string][]Question {
	m := make(map[string][]Question)

	for _, q := range questions {
		for _, t := range q.Tags {
			m[t] = append(m[t], q)
		}
	}

	return m
}

func BasedOnWorstTag(options RecommendationOptions) []Question {
	questions, err := QuestionDB.GetAny(GetQuestionOptions{})
	if err != nil {
		return questions
	}

	questionsByTag := groupByTag(questions)
	scoreByTag := make(map[string]float64)

	for t, qs := range questionsByTag {
		correct := 0.0
		total := 0.0
		for _, question := range qs {
			answers, err := AnswerDB.GetQuestionAnswered(question.Id)
			if err != nil {
				return questions
			}
			for _, answer := range answers {
				total++
				if answer.Correct {
					correct++
				}
			}
		}

		scoreByTag[t] = correct / total
	}

	lowest := 1.0
	lowestScoreTag := ""
	for tag, score := range scoreByTag {
		if score < lowest {
			lowest = score
			lowestScoreTag = tag
		}
	}

	selected := make([]Question, 0)
	lowestScoreQuestions := questionsByTag[lowestScoreTag]

	if options.limit <= 0 {
		options.limit = 1
	}
	if options.limit > len(lowestScoreQuestions) {
		options.limit = len(lowestScoreQuestions) - 1
	}

	for i := 0; i < options.limit; i++ {
		selected = append(selected, lowestScoreQuestions[rand.Intn(len(lowestScoreQuestions))])
	}

	return selected
}
