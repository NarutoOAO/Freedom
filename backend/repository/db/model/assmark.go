package model

import (
	"github.com/jinzhu/gorm"
)

// AssMark this model for AssMark
type AssMark struct {
	gorm.Model
	AssignmentId uint
	AssName      string
	StudentId    uint
	CourseNumber int
	Mark         float64
	Content      string
	FileUrl      string
	MaxScore     float64
	GroupId      uint
	GroupName    string
}

func (assMark *AssMark) FileURL() string {
	signedGetURL := assMark.FileUrl
	return signedGetURL
}
