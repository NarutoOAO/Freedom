package model

import "github.com/jinzhu/gorm"

// Forum this model for creating Forum
type Forum struct {
	gorm.Model
	CourseNumber int
	ForumName    string
	Introduction string
}
