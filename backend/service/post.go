package service

import (
	"9900project/pkg/e"
	"9900project/repository/cache"
	dao2 "9900project/repository/db/dao"
	"9900project/repository/db/model"
	"9900project/serializar"
	"context"
	"fmt"
	"strconv"
)

type PostService struct {
	ForumID int    `json:"forum_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  int    `json:"status"`
}

type VoteData struct {
	PostID    string  `json:"post_id"`
	Direction float64 `json:"direction"`
}

type SearchPostService struct {
	Info string `json:"info"`
}

func (service *PostService) CreatePost(ctx context.Context, id uint) serializar.Response {
	code := e.SUCCESS
	var post *model.Post
	dao := dao2.NewPostDao(ctx)
	dao1 := dao2.NewUserDao(ctx)
	dao3 := dao2.NewForumDao(ctx)
	forum, _ := dao3.GetForumById(uint(service.ForumID))

	user, _ := dao1.GetUserById(id)
	fmt.Println(uint(service.ForumID))
	post = &model.Post{
		ForumID:      uint(service.ForumID),
		ForumName:    forum.ForumName,
		Title:        service.Title,
		Content:      service.Content,
		AuthorId:     id,
		AuthorName:   user.NickName,
		CourseNumber: forum.CourseNumber,
	}
	post.Status = user.Authority
	err := dao.CreatePost(post)
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    "database failed",
		}
	}
	dao5 := dao2.NewCourseDao(ctx)
	course, _ := dao5.GetCourseByCourseNumber(post.CourseNumber)

	var notification *model.Notification
	notification = &model.Notification{
		Title:             service.Title,
		CourseNumber:      forum.CourseNumber,
		Status:            0,
		PostAuthorId:      id,
		PostAuthorName:    user.NickName,
		PostId:            post.ID,
		Authority:         user.Authority,
		CourseTeacherId:   course.TeacherId,
		CourseTeacherName: course.TeacherName,
	}
	dao4 := dao2.NewNotificationDao(ctx)
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
	HandleMessages(notifications, course.TeacherId)
	//err = cache.CreatePost(fmt.Sprint(post.ID), fmt.Sprint(post.AuthorId), post.Title, post.Content, fmt.Sprint(post.ForumID))
	return serializar.Response{
		Status: code,
		Msg:    "create success",
		Data:   post,
	}
}

func (service *PostService) GetPostsByForumId(ctx context.Context, fId uint) serializar.Response {
	code := e.SUCCESS
	var posts []*model.Post
	var err error
	dao := dao2.NewPostDao(ctx)
	posts, err = dao.GetPostsByForumId(fId)
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    "database failed",
		}
	}
	return serializar.Response{
		Status: code,
		Msg:    "enquire success",
		Data:   serializar.BuildPosts(posts),
	}
}

func (service *PostService) GetPostsByCourseNumber(ctx context.Context, courseNumber int) serializar.Response {
	code := e.SUCCESS
	var posts []*model.Post
	var err error
	dao := dao2.NewPostDao(ctx)
	posts, err = dao.GetPostsByCourseNumber(courseNumber)
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    "database failed",
		}
	}
	return serializar.Response{
		Status: code,
		Msg:    "enquire success",
		Data:   serializar.BuildPosts(posts),
	}
}

func (service *PostService) GetPostInformationByForumId(ctx context.Context, pId uint) serializar.Response {
	code := e.SUCCESS
	var post *model.Post
	var err error
	dao := dao2.NewPostDao(ctx)
	post, err = dao.GetPostById(pId)

	//post.AddView()

	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    "database failed",
		}
	}
	return serializar.Response{
		Status: code,
		Msg:    "enquire success",
		Data:   serializar.BuildPost(post),
	}
}

func (vote *VoteData) VoteToPost(ctx context.Context, userId uint) serializar.Response {
	code := e.SUCCESS
	err := cache.PostVote(vote.PostID, fmt.Sprint(userId), vote.Direction)
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    "redis failed",
		}
	}
	dao := dao2.NewPostDao(ctx)
	pId, _ := strconv.Atoi(vote.PostID)
	post, err := dao.GetPostById(uint(pId))
	return serializar.Response{
		Status: code,
		Msg:    "enquire success",
		Data:   serializar.BuildPost(post),
	}
}

func (service *PostService) SearchPosts(ctx context.Context, info string, courseNumber int) serializar.Response {
	code := e.SUCCESS
	var posts []*model.Post
	var err error
	dao := dao2.NewPostDao(ctx)
	posts, err = dao.SearchPostsByInfo(info, courseNumber)
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    "database failed",
		}
	}
	count := len(posts)
	return serializar.BuildListResponse(serializar.BuildPosts(posts), uint(count))
}
