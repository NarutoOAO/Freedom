package model

import "github.com/jinzhu/gorm"

type Forum struct {
	gorm.Model
	CourseNumber int
	ForumName    string
	Introduction string
}
