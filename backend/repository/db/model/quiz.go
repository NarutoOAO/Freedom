package model

import "github.com/jinzhu/gorm"

// Quiz this model is to create quiz
type Quiz struct {
	gorm.Model
	QuizName     string
	CourseNumber int
	MaxScore     float64
	StartTime    LocalTime
	EndTime      LocalTime
}
