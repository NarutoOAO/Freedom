package serializar

import (
	"9900project/conf"
	"9900project/repository/db/model"
)

type Course struct {
	CourseNumber int
	CourseName   string
	TeacherId    uint
	TeacherName  string
	CourseImg    string
}

func BuildCourse(course *model.Course) *Course {
	return &Course{
		CourseNumber: course.CourseNumber,
		CourseName:   course.CourseName,
		TeacherId:    course.TeacherId,
		TeacherName:  course.TeacherName,
		CourseImg:    conf.PhotoHost + conf.HttpPort + conf.CourseImgPath + course.ImgURL(),
	}
}

func BuildCourses(items []*model.Course) (courses []*Course) {
	for _, item := range items {
		course := BuildCourse(item)
		courses = append(courses, course)
	}
	return
}
