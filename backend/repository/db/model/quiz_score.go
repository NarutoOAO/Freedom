package model

import "github.com/jinzhu/gorm"

// QuizScore this model is to review the QuizScore
type QuizScore struct {
	gorm.Model
	QuizId         uint
	QuizQuestionId uint
	UserId         uint
	Answer         string
	Score          float64
}
