package service

import (
	"9900project/pkg/e"
	dao2 "9900project/repository/db/dao"
	"9900project/repository/db/model"
	"9900project/serializar"
	"context"
	"fmt"
)

type CommentService struct {
	Content string `json:"content"`
	PostId  uint   `json:"post_id"`
}

func (service *CommentService) CreateComment(ctx context.Context, id uint) serializar.Response {
	code := e.SUCCESS
	var comment *model.Comment
	var err error
	var user *model.User
	dao := dao2.NewCommentDao(ctx)
	dao1 := dao2.NewUserDao(ctx)
	user, _ = dao1.GetUserById(id)
	comment = &model.Comment{
		PostId:     service.PostId,
		AuthorId:   id,
		Content:    service.Content,
		AuthorName: user.NickName,
		Authority:  user.Authority,
	}
	err = dao.CreateComment(comment)
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    "database error",
			Error:  err.Error(),
		}
	}
	dao3 := dao2.NewPostDao(ctx)
	dao4 := dao2.NewNotificationDao(ctx)
	dao5 := dao2.NewCourseDao(ctx)
	post, _ := dao3.GetPostById(service.PostId)
	course, _ := dao5.GetCourseByCourseNumber(post.CourseNumber)
	fmt.Println("zhejiekeshi", course.TeacherId, ", ", course.TeacherName)
	user2, _ := dao1.GetUserById(post.AuthorId)
	var notification *model.Notification
	notification = &model.Notification{
		Title:             post.Title,
		Content:           comment.Content,
		CourseNumber:      post.CourseNumber,
		Status:            0,
		PostAuthorId:      post.AuthorId,
		PostAuthorName:    post.AuthorName,
		PostId:            service.PostId,
		CommentAuthorId:   comment.AuthorId,
		CommentAuthorName: comment.AuthorName,
		Authority:         user2.Authority,
		CourseTeacherId:   course.TeacherId,
		CourseTeacherName: course.TeacherName,
	}
	err1 := dao4.CreateNotification(notification)
	if err1 != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    "database error",
			Error:  err.Error(),
		}
	}
	var notifications []*model.Notification
	HandleMessages(notifications, post.AuthorId)
	return serializar.Response{
		Status: code,
		Msg:    "enquiry success",
		Data:   serializar.BuildComment(comment),
	}
}

func (service *CommentService) GetCommentByPostId(ctx context.Context, pId uint) serializar.Response {
	code := e.SUCCESS
	dao := dao2.NewCommentDao(ctx)
	comments, err := dao.GetCommentByPostId(pId)
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
		Msg:    "enquire success",
		Data:   serializar.BuildComments(comments),
	}
}
