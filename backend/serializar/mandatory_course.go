package serializar

import (
	"9900project/conf"
	"9900project/repository/db/model"
)

type MandatoryCourse struct {
	CourseNumber   int
	CourseName     string
	TeacherId      uint
	TeacherName    string
	CourseImg      string
	ClassTime      string
	CourseLocation string
	MaxPeople      int
	Term           string
}

func BuildMandatoryCourse(course *model.MandatoryCourse) *MandatoryCourse {
	return &MandatoryCourse{
		CourseNumber:   course.CourseNumber,
		CourseName:     course.CourseName,
		TeacherId:      course.TeacherId,
		TeacherName:    course.TeacherName,
		ClassTime:      course.ClassTime,
		CourseLocation: course.CourseLocation,
		CourseImg:      conf.PhotoHost + conf.HttpPort + conf.CourseImgPath + course.ImgURL(),
		MaxPeople:      course.MaxPeople,
		Term:           course.Term,
	}
}

func BuildMandatoryCourses(items []*model.MandatoryCourse) (courses []*MandatoryCourse) {
	for _, item := range items {
		course := BuildMandatoryCourse(item)
		courses = append(courses, course)
	}
	return
}
