package model

import "github.com/jinzhu/gorm"

type QuizQuestion struct {
	gorm.Model
	QuizId         uint
	QuestionNumber int
	Type           int
	SelectA        string
	SelectB        string
	SelectC        string
	SelectD        string
	Score          float64
	Description    string
	Answer         string
}
