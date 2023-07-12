package model

import (
	"github.com/jinzhu/gorm"
)

type AssMark struct {
	gorm.Model
	AssignmentId uint
	StudentId    uint
	CourseNumber int
	Mark         float64
	Content      string
	FileUrl      string
	MaxScore     float64
}

func (assMark *AssMark) FileURL() string {
	signedGetURL := assMark.FileUrl
	return signedGetURL
}
