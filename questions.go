package main

type QuestionType string

const (
	SingleSelectQuestionType QuestionType = "singleselect"
	MultiSelectQuestionType  QuestionType = "multiselect"
	FreeFormQuestionType     QuestionType = "freeform"
)

type Question struct {
	QuestionType         QuestionType
	SingleSelectQuestion *SingleSelectQuestion
	MultiSelectQuestion  *MultiSelectQuestion
	FreeFormQuestion     *FreeFormQuestion
}

type SingleSelectQuestion struct {
	Question string
	Answers  []string
}

type MultiSelectQuestion struct {
	Question string
	Answers  []string
}

type FreeFormQuestion struct {
	Question string
}

func getQuestions() []Question {
	return []Question{
		{
			QuestionType: FreeFormQuestionType,
			FreeFormQuestion: &FreeFormQuestion{
				Question: "What is your favorite color?",
			},
		},
		{
			QuestionType: SingleSelectQuestionType,
			SingleSelectQuestion: &SingleSelectQuestion{
				Question: "What is your favorite programming language?",
				Answers: []string{
					"Go",
					"Python",
					"JavaScript",
					"Java",
					"C#",
				},
			},
		},
	}
}
