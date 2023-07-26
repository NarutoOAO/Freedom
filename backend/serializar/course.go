package serializar

import (
	"9900project/conf"
	"9900project/repository/db/model"
)

type Course struct {
	CourseNumber   int
	CourseName     string
	TeacherId      uint
	TeacherName    string
	CourseImg      string
	ClassTime      string
	CourseLocation string
	MaxPeople      int
	Classification string
	CurrentPeople  int
}

func BuildCourse(course *model.Course, coursesSelect []*model.CourseSelect) *Course {
	return &Course{
		CourseNumber:   course.CourseNumber,
		CourseName:     course.CourseName,
		TeacherId:      course.TeacherId,
		TeacherName:    course.TeacherName,
		ClassTime:      course.ClassTime,
		CourseLocation: course.CourseLocation,
		MaxPeople:      course.MaxPeople,
		Classification: course.Classification,
		CurrentPeople:  len(coursesSelect),
		CourseImg:      conf.PhotoHost + conf.HttpPort + conf.CourseImgPath + course.ImgURL(),
	}
}

func BuildCourses(items []*model.Course) (courses []*Course) {
	for _, item := range items {
		course := BuildCourse(item, nil)
		courses = append(courses, course)
	}
	return
}
