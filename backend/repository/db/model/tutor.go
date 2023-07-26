package model

import (
	"github.com/jinzhu/gorm"
)

type Tutor struct {
	gorm.Model
	Email        string
	NickName     string
	Authority    int
	CourseNumber int
}
