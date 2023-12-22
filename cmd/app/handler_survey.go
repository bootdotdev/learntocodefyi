package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/bootdotdev/learntocodefyi/internal/sqlcdb"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

const surveyID2024 = "2fc1ba1c-ad07-49d4-936b-4b89c733c001"

func (hs *handlerState) handlerGetSurvey(w http.ResponseWriter, r *http.Request, user sqlcdb.User) {
	nextQuestion, done, err := hs.getNextQuestion(r, user)
	if err != nil {
		log.Println("Error getting next question:", err)
		http.Error(w, "error getting next question", http.StatusInternalServerError)
		return
	}
	if done {
		http.Redirect(w, r, "/thanks", http.StatusFound)
		return
	}

	hs.templates.ExecuteTemplate(w, "survey.html", nextQuestion)
}

func (hs *handlerState) getNextQuestion(r *http.Request, user sqlcdb.User) (question Question, done bool, err error) {
	responses, err := hs.queries.GetAnsweredQuestions(r.Context(), sqlcdb.GetAnsweredQuestionsParams{
		UserID:   user.ID,
		SurveyID: surveyID2024,
	})
	if err != nil {
		return Question{}, false, fmt.Errorf("error getting answered questions: %w", err)
	}
	questionIDsAnswered := make(map[string]struct{})
	for _, response := range responses {
		questionIDsAnswered[response.QuestionID] = struct{}{}
	}

	nextIndex := -1
	questions := getQuestions()
	for i, question := range questions {
		if _, ok := questionIDsAnswered[question.ID]; !ok {
			nextIndex = i
			break
		}
	}
	if nextIndex == -1 {
		return Question{}, true, nil
	}
	return questions[nextIndex], false, nil
}

func (hs *handlerState) handlerPostSurvey(w http.ResponseWriter, r *http.Request, user sqlcdb.User) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form: "+err.Error(), http.StatusBadRequest)
		return
	}
	answers := r.Form["answer"]
	finalAnswer := strings.Join(answers, "|")
	if finalAnswer == "" {
		http.Error(w, "Answer is required", http.StatusBadRequest)
		return
	}

	questionID := chi.URLParam(r, "question_id")
	if questionID == "" {
		http.Error(w, "Question ID is required", http.StatusBadRequest)
		return
	}

	questions := getQuestions()
	found := false
	for _, question := range questions {
		if question.ID == questionID {
			found = true
			break
		}
	}
	if !found {
		http.Error(w, "Question ID not found", http.StatusBadRequest)
		return
	}

	_, err := hs.queries.UpsertResponse(
		r.Context(),
		sqlcdb.UpsertResponseParams{
			ID:         uuid.New().String(),
			UserID:     user.ID,
			SurveyID:   surveyID2024,
			QuestionID: questionID,
			Answer:     finalAnswer,
		},
	)
	if err != nil {
		log.Println("Error upserting response:", err)
		http.Error(w, "error upserting response", http.StatusInternalServerError)
		return
	}

	nextQuestion, done, err := hs.getNextQuestion(r, user)
	if err != nil {
		log.Println("Error getting next question:", err)
		http.Error(w, "error getting next question", http.StatusInternalServerError)
		return
	}
	if done {
		http.Redirect(w, r, "/thanks", http.StatusFound)
		return
	}

	switch nextQuestion.QuestionType {
	case MultiSelectQuestionType:
		hs.templates.ExecuteTemplate(w, "multiselect.html", nextQuestion)
		return
	case CountryQuestionType:
		hs.templates.ExecuteTemplate(w, "country.html", nextQuestion)
		return
	case SingleSelectQuestionType:
		hs.templates.ExecuteTemplate(w, "singleselect.html", nextQuestion)
		return
	case OrderingQuestionType:
		hs.templates.ExecuteTemplate(w, "ordering.html", nextQuestion)
		return
	default:
		log.Println("Unknown question type:", nextQuestion.QuestionType)
		http.Error(w, "unknown question type", http.StatusInternalServerError)
		return
	}
}
