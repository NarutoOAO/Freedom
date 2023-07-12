package model

import "gorm.io/gorm"

type CourseSelect struct {
	gorm.Model
	CourseNumber int
	CourseName   string
	CourseImg    string
	TeacherId    uint
	TeacherName  string
	StudentId    uint
	Status       int
}

func (courseSelect *CourseSelect) CourseImgURL() string {
	url := courseSelect.CourseImg
	return url
}
