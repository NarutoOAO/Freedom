package service

import (
	"9900project/pkg/e"
	util "9900project/pkg/utils"
	dao2 "9900project/repository/db/dao"
	"9900project/repository/db/model"
	"9900project/serializar"
	"context"
	"fmt"
	"mime/multipart"
)

type AssignmentService struct {
	AssignmentId int             `form:"assignment_id" json:"assignment_id"`
	CourseNumber int             `form:"course_number" json:"course_number"`
	FileName     string          `form:"file_name" json:"file_name"`
	EndTime      model.LocalTime `form:"end_time" json:"end_time"`
	MaxScore     float64         `form:"max_score" json:"max_score"`
}

func (service *AssignmentService) UploadAssignment(ctx context.Context, file multipart.File, fileHeader int64) serializar.Response {
	code := e.SUCCESS
	var err error
	dao := dao2.NewAssignmentDao(ctx)
	assignment := &model.Assignment{
		CourseNumber: service.CourseNumber,
		FileName:     service.FileName,
		EndTime:      service.EndTime,
		MaxScore:     service.MaxScore,
	}
	var path string
	path, err = util.UploadAssignmentToLocalStatic(file, service.CourseNumber, service.FileName)
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	assignment.FileUrl = path
	err = dao.CreateAssignment(assignment)
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
		Data:   serializar.BuildAssignment(assignment),
		Msg:    "insert success",
	}
}

func (service *AssignmentService) UpdateAssignment(ctx context.Context) serializar.Response {
	fmt.Println(service.EndTime)
	code := e.SUCCESS
	var err error
	dao := dao2.NewAssignmentDao(ctx)
	assignment, err := dao.GetAssignmentById(uint(service.AssignmentId))
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	assignment.EndTime = service.EndTime
	err = dao.UpdateAssignment(uint(service.AssignmentId), assignment)
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
		Data:   serializar.BuildAssignment(assignment),
		Msg:    "update success",
	}
}

func (service *AssignmentService) GetAssignments(ctx context.Context, courseNumber int) serializar.Response {
	code := e.SUCCESS
	var err error
	var assignments []*model.Assignment
	dao := dao2.NewAssignmentDao(ctx)
	assignments, err = dao.GetAllAssignments(courseNumber)
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
		Data:   serializar.BuildAssignments(assignments),
		Msg:    "enquiry success",
	}
}

func (service *AssignmentService) DeleteAssignment(ctx context.Context) serializar.Response {
	code := e.SUCCESS
	var err error
	dao := dao2.NewAssignmentDao(ctx)
	err = dao.DeleteAssignment(uint(service.AssignmentId))
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
