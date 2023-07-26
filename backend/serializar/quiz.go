package serializar

import (
	"9900project/repository/db/model"
)

type Quiz struct {
	Id           uint            `json:"quiz_id"`
	CourseNumber int             `json:"course_number"`
	QuizName     string          `json:"quiz_name"`
	MaxScore     float64         `json:"max_score"`
	StartTime    model.LocalTime `json:"start_time"`
	EndTime      model.LocalTime `json:"end_time"`
}

func BuildQuiz(quiz *model.Quiz) *Quiz {
	return &Quiz{
		Id:           quiz.ID,
		CourseNumber: quiz.CourseNumber,
		QuizName:     quiz.QuizName,
		MaxScore:     quiz.MaxScore,
		StartTime:    quiz.StartTime,
		EndTime:      quiz.EndTime,
	}
}

func BuildQuizs(items []*model.Quiz) (quizs []*Quiz) {
	for _, item := range items {
		quiz := BuildQuiz(item)
		quizs = append(quizs, quiz)
	}
	return
}
