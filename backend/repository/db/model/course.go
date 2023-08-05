package model

import (
	"gorm.io/gorm"
)

// Course this model is creating Course
type Course struct {
	gorm.Model
	CourseNumber   int `gorm:"unique"`
	CourseName     string
	TeacherId      uint
	TeacherName    string
	CourseImg      string
	ClassTime      string
	CourseLocation string
	MaxPeople      int
	Classification string
	Term           int
}

func (course *Course) ImgURL() string {
	signedGetURL := course.CourseImg
	return signedGetURL
}
