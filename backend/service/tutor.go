package service

import (
	"9900project/pkg/e"
	dao2 "9900project/repository/db/dao"
	"9900project/repository/db/model"
	"9900project/serializar"
	"context"
)

// CreateTutorService is a struct to create tutor
type CreateTutorService struct {
	UserId       uint   `json:"user_id"`
	Email        string `json:"email"`
	NickName     string `json:"nick_name"`
	Authority    int    `json:"authority"`
	CourseNumber int    `json:"course_number"`
}

// define a service to get tutor
type GetTutorService struct {
	ID           uint   `json:"id"`
	UserId       uint   `json:"user_id"`
	Email        string `json:"email"`
	NickName     string `json:"nick_name"`
	Authority    int    `json:"authority"`
	CourseNumber int    `json:"course_number"`
}

// create tutor
func (service *CreateTutorService) CreateTutor(ctx context.Context) serializar.Response {
	code := e.SUCCESS
	var err error
	dao := dao2.NewTutorDao(ctx)
	_, exist, err := dao.IfExistOrNot(service.UserId, service.CourseNumber)
	if exist {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    "Tutor is existed",
		}
	}
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	tutor := &model.Tutor{
		UserId:       service.UserId,
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

// get tutors by course number
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

// delete tutor by id
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
