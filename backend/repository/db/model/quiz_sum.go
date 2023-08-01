package model

import "github.com/jinzhu/gorm"

type QuizSum struct {
	gorm.Model
	QuizId       uint
	QuizName     string
	CourseNumber int
	UserId       uint
	Score        float64
	MaxScore     float64
}
