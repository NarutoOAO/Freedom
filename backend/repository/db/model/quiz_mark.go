package model

import "github.com/jinzhu/gorm"

// QuizMark this model to create quizmark
type QuizMark struct {
	gorm.Model
	QuizId          uint
	QuestionId      uint
	QuestionNumber  int
	UserId          uint
	QuizDescription string
	QuizAnswer      string
	Type            int
	UserAnswer      string
	Score           float64
	MaxScore        float64
}
