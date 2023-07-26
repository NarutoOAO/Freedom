package serializar

import (
	"9900project/repository/db/model"
)

type QuizQuestion struct {
	QuizId         uint    `json:"quiz_id"`
	QuestionNumber int     `json:"question_number"`
	Type           int     `json:"type"` //1为单选，2为多选，3为简答
	Score          float64 `json:"score"`
	SelectA        string  `json:"select_A"`
	SelectB        string  `json:"select_B"`
	SelectC        string  `json:"select_C"`
	SelectD        string  `json:"select_D"`
	Description    string  `json:"description"`
	Answer         string  `json:"answer"`
}

func BuildQuizQuestion(quizQuestion *model.QuizQuestion) *QuizQuestion {
	return &QuizQuestion{
		QuizId:         quizQuestion.QuizId,
		QuestionNumber: quizQuestion.QuestionNumber,
		Type:           quizQuestion.Type,
		SelectA:        quizQuestion.SelectA,
		SelectB:        quizQuestion.SelectB,
		SelectC:        quizQuestion.SelectC,
		SelectD:        quizQuestion.SelectD,
		Score:          quizQuestion.Score,
		Description:    quizQuestion.Description,
		Answer:         quizQuestion.Answer,
	}
}

func BuildQuizQuestions(items []*model.QuizQuestion) (quizQuestions []*QuizQuestion) {
	for _, item := range items {
		quizQuestion := BuildQuizQuestion(item)
		quizQuestions = append(quizQuestions, quizQuestion)
	}
	return
}
