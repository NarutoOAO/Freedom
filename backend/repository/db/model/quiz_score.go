package model

import "github.com/jinzhu/gorm"

type QuizScore struct {
	gorm.Model
	QuizId         uint
	QuizQuestionId uint
	UserId         uint
	Answer         string
	Score          float64
}
