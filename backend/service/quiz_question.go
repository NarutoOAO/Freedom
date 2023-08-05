package service

import (
	"9900project/pkg/e"
	dao2 "9900project/repository/db/dao"
	"9900project/repository/db/model"
	"9900project/serializar"
	"context"
)

type QuizQuestionService struct {
	QuizId         uint    `json:"quiz_id"`
	QuestionNumber int     `json:"question_number"`
	Type           int     `json:"type"` //1为单选，2为多选，3为简答
	Score          float64 `json:"score"`
	SelectA        string  `json:"select_A"`
	SelectB        string  `json:"select_B"`
	SelectC        string  `json:"select_C"`
	SelectD        string  `json:"select_D"`
	Description    string  `json:"description"`
	Answer         string  `json:"answer"`
}

func (service *QuizQuestionService) CreateQuizQuestion(ctx context.Context) serializar.Response {
	code := e.SUCCESS
	var err error
	dao := dao2.NewQuizQuestionDao(ctx)
	quizQuestion := &model.QuizQuestion{
		QuizId:         service.QuizId,
		QuestionNumber: service.QuestionNumber,
		Type:           service.Type,
		SelectA:        service.SelectA,
		SelectB:        service.SelectB,
		SelectC:        service.SelectC,
		SelectD:        service.SelectD,
		Score:          service.Score,
		Description:    service.Description,
		Answer:         service.Answer,
	}
	err = dao.CreateQuizQuestion(quizQuestion)
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
		Data:   serializar.BuildQuizQuestion(quizQuestion),
		Msg:    "insert success",
	}
}

func (service *QuizQuestionService) GetQuizQuestions(ctx context.Context, quizId uint) serializar.Response {
	code := e.SUCCESS
	var err error
	var quizQuestions []*model.QuizQuestion
	dao := dao2.NewQuizQuestionDao(ctx)
	quizQuestions, err = dao.GetQuizQuestionsByQuizId(quizId)
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
		Data:   serializar.BuildQuizQuestions(quizQuestions),
		Msg:    "enquiry success",
	}
}
