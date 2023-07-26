package service

import (
	"9900project/pkg/e"
	dao2 "9900project/repository/db/dao"
	"9900project/repository/db/model"
	"9900project/serializar"
	"context"
	"fmt"
)

type CourseSelect struct {
	CourseNumber   int    `json:"course_number"`
	Classification string `json:"classification"`
}

// student select course
func (service *CourseSelect) SelectCourse(ctx context.Context, id uint) serializar.Response {
	code := e.SUCCESS
	var err error
	var course *model.Course
	dao := dao2.NewCourseSelectDao(ctx)
	dao1 := dao2.NewCourseDao(ctx)
	course, err = dao1.GetCourseByCourseNumber(service.CourseNumber)
	if course == nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    "course is not existed",
			Error:  err.Error(),
		}
	}
	courses, err := dao.GetCourseByCourseNumber(service.CourseNumber)
	if courses != nil && len(courses) >= course.MaxPeople {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    "The maximum enrollment limit is exceeded",
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

func (service *CourseSelect) GetCoursesByNumber(ctx context.Context) interface{} {
	code := e.SUCCESS
	var err error
	//var courseSelect []*model.CourseSelect
	dao := dao2.NewCourseSelectDao(ctx)
	dao1 := dao2.NewCourseDao(ctx)
	var course *model.Course
	fmt.Printf("", service)
	fmt.Printf("", service.CourseNumber)
	fmt.Printf("", dao1)
	course, _ = dao1.GetCourseByCourseNumber(service.CourseNumber)
	if course == nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    "course is not existed",
			Error:  err.Error(),
		}
	}
	courses, _ := dao.GetCourseByCourseNumber(service.CourseNumber)
	return serializar.Response{
		Status: code,
		Msg:    "success",
		Data:   serializar.BuildCourse(course, courses),
	}
}

func (service *CourseSelect) StudentDropCourse(ctx context.Context, CourseNumber int, id uint) serializar.Response {
	code := e.SUCCESS
	var err error
	dao := dao2.NewCourseSelectDao(ctx)
	err = dao.DropCourseById(CourseNumber, id)
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    "Drop course failed",
			Error:  err.Error(),
		}
	}
	return serializar.Response{
		Status: code,
		Msg:    "Drop course success",
	}
}

func (service *CourseSelect) Statistics(ctx context.Context, id uint) serializar.Response {
	code := e.SUCCESS
	var err error
	var courseSelect []*model.CourseSelect
	dao := dao2.NewCourseSelectDao(ctx)
	dao1 := dao2.NewCourseDao(ctx)
	//查询已选课程数量
	fmt.Printf("", service.Classification)
	courseSelect, err = dao.GetCourseByClassification(id, service.Classification)
	fmt.Printf("", err)
	fmt.Printf("", courseSelect)
	//查询未选择课程
	data := []int{}
	for _, value := range courseSelect {
		data = append(data, value.CourseNumber)
	}
	courses, err := dao1.GetByNotSelected(service.Classification, data)
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
		Data: &struct {
			Credit  int
			Courses []*model.Course
		}{
			Credit:  len(courseSelect) * 6,
			Courses: courses,
		},
	}
}
