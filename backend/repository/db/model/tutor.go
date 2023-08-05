package model

import (
	"github.com/jinzhu/gorm"
)

// Tutor this model is to create tutor group.
type Tutor struct {
	gorm.Model
	UserId       uint
	Email        string
	NickName     string
	Authority    int
	CourseNumber int
}
