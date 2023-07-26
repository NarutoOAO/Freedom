package service

import (
	"9900project/conf"
	"9900project/pkg/e"
	dao2 "9900project/repository/db/dao"
	"9900project/repository/db/model"
	"9900project/serializar"
	"context"
)

type CourseService struct {
	CourseNumber   int    `json:"course_number"`
	CourseName     string `json:"course_name"`
	TeacherId      uint   `json:"teacher_id"`
	ClassTime      string `json:"class_time" `
	CourseLocation string `json:"course_location"`
	MaxPeople      int    `json:"max_people"`
	Classification string `json:"Classification"`
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
	if service.MaxPeople == 0 {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    "The maximum number of applicants cannot be 0",
			Error:  err.Error(),
		}
	}
	if service.Classification == "" {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    "Course categories cannot be empty",
			Error:  err.Error(),
		}
	}
	course = &model.Course{
		CourseNumber:   service.CourseNumber,
		CourseName:     service.CourseName,
		ClassTime:      service.ClassTime,
		CourseLocation: service.CourseLocation,
		MaxPeople:      service.MaxPeople,
		Classification: service.Classification,
		TeacherId:      id,
		TeacherName:    user.NickName,
		CourseImg:      "course.JPG",
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
		Data:   serializar.BuildCourse(course, nil),
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

func (service *CourseSelect) StudentSelectCourse(ctx context.Context, id uint) serializar.Response {
	code := e.SUCCESS
	var err error
	var courses []*model.Course
	dao := dao2.NewCourseDao(ctx)
	dao1 := dao2.NewCourseSelectDao(ctx)
	courses, err = dao.StudentSelectCourse(id)
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    "database error",
			Error:  err.Error(),
		}
	}
	var scourse []*serializar.Course
	for _, item := range courses {
		coursesSelect, _ := dao1.GetCourseByCourseNumber(item.CourseNumber)
		var couse = &serializar.Course{
			CourseNumber:   item.CourseNumber,
			CourseName:     item.CourseName,
			TeacherId:      item.TeacherId,
			TeacherName:    item.TeacherName,
			ClassTime:      item.ClassTime,
			CourseLocation: item.CourseLocation,
			MaxPeople:      item.MaxPeople,
			Classification: item.Classification,
			CurrentPeople:  len(coursesSelect),
			CourseImg:      conf.PhotoHost + conf.HttpPort + conf.CourseImgPath + item.ImgURL(),
		}
		scourse = append(scourse, couse)
	}

	return serializar.Response{
		Status: code,
		Msg:    "success",
		Data:   scourse,
	}
}
