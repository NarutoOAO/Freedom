package model

import "github.com/jinzhu/gorm"

// Score this model to review the grade
type Score struct {
	gorm.Model
	UserId       uint
	CourseNumber int
	Score        float64
	MaxScore     float64
}
