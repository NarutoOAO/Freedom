package service

import (
	"9900project/pkg/e"
	dao2 "9900project/repository/db/dao"
	"9900project/repository/db/model"
	"9900project/serializar"
	"context"
)

type QuizService struct {
	QuizId       int             `json:"quiz_id"`
	QuizName     string          `json:"quiz_name"`
	CourseNumber int             `json:"course_number"`
	MaxScore     float64         `json:"max_score"`
	StartTime    model.LocalTime `json:"start_time"`
	EndTime      model.LocalTime `json:"end_time"`
}

func (service *QuizService) CreateQuiz(ctx context.Context) serializar.Response {
	code := e.SUCCESS
	var err error
	dao := dao2.NewQuizDao(ctx)
	quiz := &model.Quiz{
		QuizName:     service.QuizName,
		CourseNumber: service.CourseNumber,
		MaxScore:     service.MaxScore,
		StartTime:    service.StartTime,
		EndTime:      service.EndTime,
	}
	err = dao.CreateQuiz(quiz)
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
		Data:   serializar.BuildQuiz(quiz),
		Msg:    "insert success",
	}
}

func (service *QuizService) GetQuizs(ctx context.Context, courseNumber int) serializar.Response {
	code := e.SUCCESS
	var err error
	var quizs []*model.Quiz
	dao := dao2.NewQuizDao(ctx)
	quizs, err = dao.GetAllQuiz(courseNumber)
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
		Data:   serializar.BuildQuizs(quizs),
		Msg:    "enquiry success",
	}
}
