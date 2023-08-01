package model

import (
	"gorm.io/gorm"
)

type MandatoryCourse struct {
	gorm.Model
	CourseNumber   int
	CourseName     string
	TeacherId      uint
	TeacherName    string
	CourseImg      string
	ClassTime      string
	CourseLocation string
	Classification string
	MaxPeople      int
	Term           string
}

func (course *MandatoryCourse) ImgURL() string {
	signedGetURL := course.CourseImg
	return signedGetURL
}
