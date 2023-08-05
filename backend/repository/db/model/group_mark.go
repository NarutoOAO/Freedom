package model

import (
	"gorm.io/gorm"
)

// GroupMark this model is creating GroupMark
type GroupMark struct {
	gorm.Model
	GroupId         uint
	GroupName       string
	CourseNumber    int
	TeacherId       uint
	TeacherName     string
	ResponsibleId   uint
	ResponsibleName string
	AssignmentId    uint
	AssMarkId       uint
}
