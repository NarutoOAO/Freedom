package service

import (
	"9900project/pkg/e"
	dao2 "9900project/repository/db/dao"
	"9900project/repository/db/model"
	"9900project/serializar"
	"context"
)

type CreateTutorService struct {
	Email        string `json:"email"`
	NickName     string `json:"nick_name"`
	Authority    int    `json:"authority"`
	CourseNumber int    `json:"course_number"`
}

type GetTutorService struct {
	ID           uint   `json:"id"`
	Email        string `json:"email"`
	NickName     string `json:"nick_name"`
	Authority    int    `json:"authority"`
	CourseNumber int    `json:"course_number"`
}

func (service *CreateTutorService) CreateTutor(ctx context.Context) serializar.Response {
	code := e.SUCCESS
	var err error
	dao := dao2.NewTutorDao(ctx)
	tutor := &model.Tutor{
		Email:        service.Email,
		NickName:     service.NickName,
		Authority:    service.Authority,
		CourseNumber: service.CourseNumber,
	}
	err = dao.CreateTutor(tutor)
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
		Data:   serializar.BuildTutor(tutor),
		Msg:    "Add tutor success",
	}
}

func (service *GetTutorService) GetTutors(ctx context.Context, courseNumber int) serializar.Response {
	code := e.SUCCESS
	var err error
	var tutors []*model.Tutor
	dao := dao2.NewTutorDao(ctx)
	tutors, err = dao.GetTutors(courseNumber)
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
		Data:   serializar.BuildTutors(tutors),
		Msg:    "enquiry success",
	}
}

func (service *GetTutorService) DeleteTutorById(ctx context.Context, id uint) serializar.Response {
	code := e.SUCCESS
	var err error
	dao := dao2.NewTutorDao(ctx)
	err = dao.DeleteTutorById(id)
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
