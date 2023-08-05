package model

import "github.com/jinzhu/gorm"

// Notification this model is creating Notification
type Notification struct {
	gorm.Model
	Title             string
	Content           string
	CourseNumber      int
	Status            int
	PostAuthorId      uint
	PostAuthorName    string
	PostId            uint
	CommentAuthorId   uint
	CommentAuthorName string
	Authority         int
	CourseTeacherId   uint
	CourseTeacherName string
}
