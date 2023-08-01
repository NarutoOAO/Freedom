package model

import "github.com/jinzhu/gorm"

type Score struct {
	gorm.Model
	UserId       uint
	CourseNumber int
	Score        float64
	MaxScore     float64
}
