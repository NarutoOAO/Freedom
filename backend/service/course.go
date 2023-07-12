package service

import (
	"9900project/pkg/e"
	dao2 "9900project/repository/db/dao"
	"9900project/repository/db/model"
	"9900project/serializar"
	"context"
)

type CourseService struct {
	CourseNumber int    `json:"course_number"`
	CourseName   string `json:"course_name"`
	TeacherId    uint   `json:"teacher_id"`
}

func (service *CourseService) CreateCourse(ctx context.Context, id uint) serializar.Response {
	code := e.SUCCESS
	var err error
	var course *model.Course
	var user *model.User
	dao := dao2.NewCourseDao(ctx)
	dao1 := dao2.NewUserDao(ctx)
	user, _ = dao1.GetUserById(id)
	if service.CourseName == "" {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    "course name can not be empty",
			Error:  err.Error(),
		}
	}
	course = &model.Course{
		CourseNumber: service.CourseNumber,
		CourseName:   service.CourseName,
		TeacherId:    id,
		TeacherName:  user.NickName,
		CourseImg:    "course.JPG",
	}
	err = dao.CreateCourse(course)
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
		Msg:    "create success",
		Data:   serializar.BuildCourse(course),
	}
}

// 老师注册的课程
func (service *CourseService) ShowTeacherCourseList(ctx context.Context, id uint) serializar.Response {
	code := e.SUCCESS
	var err error
	var courses []*model.Course
	dao := dao2.NewCourseDao(ctx)
	courses, err = dao.GetCourseByTeacherId(id)
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    "enquire failed",
			Error:  err.Error(),
		}
	}
	return serializar.Response{
		Status: code,
		Msg:    "enquire success",
		Data:   serializar.BuildCourses(courses),
	}
}
