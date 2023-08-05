package serializar

import (
	"9900project/conf"
	"9900project/repository/db/model"
)

// create course select
type CourseSelect struct {
	CourseNumber int
	CourseName   string
	TeacherId    uint
	TeacherName  string
	CourseImg    string
	StudentId    uint
}

// build course select
func BuildCourseSelect(courseSelect *model.CourseSelect) *CourseSelect {
	return &CourseSelect{
		CourseNumber: courseSelect.CourseNumber,
		CourseName:   courseSelect.CourseName,
		TeacherId:    courseSelect.TeacherId,
		TeacherName:  courseSelect.TeacherName,
		StudentId:    courseSelect.StudentId,
		CourseImg:    conf.PhotoHost + conf.HttpPort + conf.CourseImgPath + courseSelect.CourseImgURL(),
	}
}

// build courses select
func BuildCoursesSelect(items []*model.CourseSelect) (coursesSelect []*CourseSelect) {
	for _, item := range items {
		courseSelect := BuildCourseSelect(item)
		coursesSelect = append(coursesSelect, courseSelect)
	}
	return

}
