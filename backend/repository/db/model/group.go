package model

import (
	"gorm.io/gorm"
)

// TutorGroup this model is creating TutorGroup
type TutorGroup struct {
	gorm.Model
	GroupName       string
	CourseNumber    int
	TeacherId       uint
	TeacherName     string
	ResponsibleId   uint
	ResponsibleName string
	AssignmentId    uint
	AssMarkId       uint
}
