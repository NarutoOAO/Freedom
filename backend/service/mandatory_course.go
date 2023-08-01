package service

import (
	"9900project/pkg/e"
	dao2 "9900project/repository/db/dao"
	"9900project/repository/db/model"
	"9900project/serializar"
	"context"
	"fmt"
)

type MandatoryCourseService struct {
	Classification string `json:"classification"`
}

// 获取分类课程
func (service *MandatoryCourseService) GetUserMandatoryCourse(ctx context.Context, id uint) serializar.Response {
	code := e.SUCCESS
	var err error
	var courses []*model.MandatoryCourse
	dao := dao2.NewUserDao(ctx)
	dao1 := dao2.NewMandatoryCourseDao(ctx)
	loginuser, err := dao.GetUserById(id)
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    "enquire failed",
			Error:  err.Error(),
		}
	}
	courses, err = dao1.GetByClassification(loginuser.Studyoption)
	fmt.Printf("", courses)
	return serializar.Response{
		Status: code,
		Msg:    "enquire success",
		Data:   serializar.BuildMandatoryCourses(courses),
	}
}
