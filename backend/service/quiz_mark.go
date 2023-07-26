package service

import (
	"9900project/pkg/e"
	dao2 "9900project/repository/db/dao"
	"9900project/repository/db/model"
	"9900project/serializar"
	"context"
	"strings"
)

type QuizMarkService struct {
	QuizQuestionId uint   `json:"quiz_question_id"`
	UserAnswer     string `json:"user_answer"`
	QuizId         uint   `json:"quiz_id"`
}

func (service *QuizMarkService) CreateQuizMark(ctx context.Context, uId uint) serializar.Response {
	code := e.SUCCESS
	var err error
	dao1 := dao2.NewQuizQuestionDao(ctx)
	dao3 := dao2.NewQuizMarkDao(ctx)
	quizQuestion, _ := dao1.GetQuizQuestionsById(service.QuizQuestionId)
	var score float64
	if quizQuestion.Type == 1 || quizQuestion.Type == 2 {
		score = ScorePart(quizQuestion.Answer, service.UserAnswer) * quizQuestion.Score
	}
	ans, _ := dao3.GetQuizMarkByPerson(service.QuizQuestionId, uId)
	if ans != nil {
		ans.UserAnswer = service.UserAnswer
		ans.Score = score
		dao3.UpdateQuizMark(ans.ID, ans)
		return serializar.Response{
			Status: code,
			Data:   serializar.BuildQuizMark(ans),
			Msg:    "update success",
		}
	}
	quizMark := &model.QuizMark{
		QuizId:          quizQuestion.QuizId,
		QuestionId:      service.QuizQuestionId,
		UserId:          uId,
		QuestionNumber:  quizQuestion.QuestionNumber,
		QuizDescription: quizQuestion.Description,
		QuizAnswer:      quizQuestion.Answer,
		Type:            quizQuestion.Type,
		UserAnswer:      service.UserAnswer,
		Score:           score,
		MaxScore:        quizQuestion.Score,
	}
	err = dao3.CreateQuizMark(quizMark)
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
		Data:   serializar.BuildQuizMark(quizMark),
		Msg:    "insert success",
	}
}

func ScorePart(a, b string) float64 {
	if a == b {
		return 1
	} else if strings.Contains(a, b) {
		return 0.5
	}
	return 0
}

func (service *QuizMarkService) GetQuizMark(ctx context.Context, uId uint) serializar.Response {
	code := e.SUCCESS
	var err error
	dao3 := dao2.NewQuizMarkDao(ctx)
	quizMarks, err := dao3.GetQuizMarksByPerson(service.QuizId, uId)
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
		Data:   serializar.BuildQuizMarks(quizMarks),
		Msg:    "insert success",
	}
}
