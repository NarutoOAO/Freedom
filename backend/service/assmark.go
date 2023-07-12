package service

import (
	"9900project/pkg/e"
	util "9900project/pkg/utils"
	dao2 "9900project/repository/db/dao"
	"9900project/repository/db/model"
	"9900project/serializar"
	"context"
	"mime/multipart"
)

type AssMarkService struct {
	AssMarkId    uint    `form:"ass_mark_id" json:"ass_mark_id"`
	AssignmentId uint    `form:"assignment_id" json:"assignment_id"`
	Mark         float64 `form:"mark" json:"mark"`
	Content      string  `form:"content" json:"content"`
	CourseNumber int     `form:"course_number" json:"course_number"`
}

func (service *AssMarkService) UploadAssSolution(ctx context.Context, file multipart.File, fileHeader int64, uId uint) serializar.Response {
	code := e.SUCCESS
	var err error
	dao := dao2.NewAssMarkDao(ctx)
	dao1 := dao2.NewUserDao(ctx)
	dao3 := dao2.NewAssignmentDao(ctx)
	user, _ := dao1.GetUserById(uId)
	assignment, _ := dao3.GetAssignmentById(service.AssignmentId)
	count, err := dao.GetAssMarkByAssId(uId, service.AssignmentId)
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	if count != 0 {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    "please delete the previous assignment",
		}
	}
	assMark := &model.AssMark{
		AssignmentId: service.AssignmentId,
		StudentId:    uId,
		CourseNumber: assignment.CourseNumber,
		MaxScore:     assignment.MaxScore,
	}
	var path string
	path, err = util.UploadAssSolutionToLocalStatic(file, assignment.CourseNumber, user.NickName)
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	assMark.FileUrl = path
	err = dao.CreateAssMark(assMark)
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializar.Response{
		Status: code,
		Data:   serializar.BuildAssMark(assMark),
		Msg:    "insert success",
	}
}

func (service *AssMarkService) DeleteAssSolution(ctx context.Context) serializar.Response {
	code := e.SUCCESS
	var err error
	dao := dao2.NewAssMarkDao(ctx)
	err = dao.DeleteAssMark(service.AssMarkId)
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializar.Response{
		Status: code,
		Msg:    "delete success",
	}
}

func (service *AssMarkService) GetAssMarks(ctx context.Context, courseNumber int, assignmentId uint) serializar.Response {
	code := e.SUCCESS
	var err error
	var assMarks []*model.AssMark
	dao := dao2.NewAssMarkDao(ctx)
	assMarks, err = dao.GetAssMarkByCourseNumber(courseNumber, assignmentId)
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializar.Response{
		Status: code,
		Data:   serializar.BuildAssMarks(assMarks),
		Msg:    "enquiry success",
	}
}

func (service *AssMarkService) GetAssMarksByStudent(ctx context.Context, uId uint, courseNumber int) serializar.Response {
	code := e.SUCCESS
	var err error
	var assMarks []*model.AssMark
	dao := dao2.NewAssMarkDao(ctx)
	assMarks, err = dao.GetAssMarkById(uId, courseNumber)
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializar.Response{
		Status: code,
		Data:   serializar.BuildAssMarks(assMarks),
		Msg:    "enquiry success",
	}
}

func (service *AssMarkService) UpdateAssMark(ctx context.Context) serializar.Response {
	code := e.SUCCESS
	var err error
	dao := dao2.NewAssMarkDao(ctx)
	assMark, err := dao.GetAssMark(service.AssMarkId)
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	assMark.Mark = service.Mark
	assMark.Content = service.Content
	err = dao.UpdateAssMark(service.AssMarkId, assMark)
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializar.Response{
		Status: code,
		Data:   serializar.BuildAssMark(assMark),
		Msg:    "update success",
	}
}
