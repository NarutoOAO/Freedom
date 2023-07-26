package model

import "github.com/jinzhu/gorm"

type Quiz struct {
	gorm.Model
	QuizName     string
	CourseNumber int
	MaxScore     float64
	StartTime    LocalTime
	EndTime      LocalTime
}
