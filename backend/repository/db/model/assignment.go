package model

import (
	"github.com/jinzhu/gorm"
)

type Assignment struct {
	gorm.Model
	EndTime      LocalTime
	FileUrl      string
	FileName     string
	CourseNumber int
	MaxScore     float64
}

func (assignment *Assignment) FileURL() string {
	signedGetURL := assignment.FileUrl
	return signedGetURL
}
