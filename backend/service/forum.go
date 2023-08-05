package service

import (
	"9900project/pkg/e"
	dao2 "9900project/repository/db/dao"
	"9900project/repository/db/model"
	"9900project/serializar"
	"context"
)

// ForumService is a struct to create forum
type ForumService struct {
	CourseNumber int    `json:"course_number"`
	ForumName    string `json:"forum_name"`
	Introduction string `json:"introduction"`
}

// CreateForum is a function to create forum
func (service *ForumService) CreateForum(ctx context.Context) serializar.Response {
	code := e.SUCCESS
	var err error
	var forum *model.Forum
	dao := dao2.NewForumDao(ctx)
	if service.ForumName == "" {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    "forum name can not be empty",
			Error:  err.Error(),
		}
	}
	forum = &model.Forum{
		CourseNumber: service.CourseNumber,
		ForumName:    service.ForumName,
		Introduction: service.Introduction,
	}
	err = dao.CreateForum(forum)
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    "create forum failed",
			Error:  err.Error(),
		}
	}
	return serializar.Response{
		Status: code,
		Msg:    "create success",
		Data:   serializar.BuildForum(forum),
	}
}

// ShowForumList is a function to show forum list
func (service *ForumService) ShowForumList(ctx context.Context, courseNumber int) serializar.Response {
	code := e.SUCCESS
	var err error
	var forums []*model.Forum
	dao := dao2.NewForumDao(ctx)
	forums, err = dao.GetForumsByCourseNumber(courseNumber)
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
		Data:   serializar.BuildForums(forums),
	}
}
