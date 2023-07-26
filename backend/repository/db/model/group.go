package model

import (
	"gorm.io/gorm"
)

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
