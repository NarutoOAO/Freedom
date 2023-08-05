package model

import "github.com/jinzhu/gorm"

// QuizSum this model is to view QuizSum
type QuizSum struct {
	gorm.Model
	QuizId       uint
	QuizName     string
	CourseNumber int
	UserId       uint
	Score        float64
	MaxScore     float64
}
