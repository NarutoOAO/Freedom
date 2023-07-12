package model

import "github.com/jinzhu/gorm"

type Quiz struct {
	gorm.Model
	QuestionId  uint
	Answer      string
	Description string
	score       float64
	Type        uint
}
