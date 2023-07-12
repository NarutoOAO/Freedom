package service

import (
	"9900project/pkg/e"
	dao2 "9900project/repository/db/dao"
	"9900project/repository/db/model"
	"9900project/serializar"
	"context"
)

type CourseSelect struct {
	CourseNumber int `json:"course_number"`
}

// student select course
func (service *CourseSelect) SelectCourse(ctx context.Context, id uint) serializar.Response {
	code := e.SUCCESS
	var err error
	var course *model.Course
	dao := dao2.NewCourseSelectDao(ctx)
	dao1 := dao2.NewCourseDao(ctx)
	course, _ = dao1.GetCourseByCourseNumber(service.CourseNumber)
	if course == nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    "course is not existed",
			Error:  err.Error(),
		}
	}
	courseSelect := &model.CourseSelect{
		CourseNumber: course.CourseNumber,
		CourseName:   course.CourseName,
		TeacherId:    course.TeacherId,
		TeacherName:  course.TeacherName,
		CourseImg:    course.CourseImg,
		StudentId:    id,
		Status:       0,
	}
	err = dao.CreateCourseSelect(courseSelect)
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    "select course failed",
			Error:  err.Error(),
		}
	}
	return serializar.Response{
		Status: code,
		Msg:    "select course success",
	}
}

// show courses
func (service *CourseSelect) GetCoursesSelectById(ctx context.Context, id uint) serializar.Response {
	code := e.SUCCESS
	var err error
	var courseSelect []*model.CourseSelect
	dao := dao2.NewCourseSelectDao(ctx)
	courseSelect, err = dao.GetCourseByStudentId(id)
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    "database error",
			Error:  err.Error(),
		}
	}
	return serializar.Response{
		Status: code,
		Msg:    "success",
		Data:   serializar.BuildCoursesSelect(courseSelect),
	}
}
