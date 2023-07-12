package model

import "github.com/jinzhu/gorm"

type Material struct {
	gorm.Model
	CourseNumber int
	FileName     string
	FileUrl      string
	FileCategory string
	Type         int //1为pdf 2为ppt
	Publish      int
}

func (material *Material) FileURL() string {
	signedGetURL := material.FileUrl
	return signedGetURL
}
