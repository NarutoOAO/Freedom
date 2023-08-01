package service

import (
	"9900project/pkg/e"
	dao2 "9900project/repository/db/dao"
	"9900project/repository/db/model"
	"9900project/serializar"
	"context"
)

type QuizSumService struct {
}

func (service *QuizSumService) GetQuizSumByStudent(ctx context.Context, userId uint, courseNumber int) serializar.Response {
	code := e.SUCCESS
	var err error
	quizSums := make([]*model.QuizSum, 0)
	dao := dao2.NewQuizDao(ctx)
	dao1 := dao2.NewQuizSumDao(ctx)
	//dao3 := dao2.NewQuizQuestionDao(ctx)
	dao4 := dao2.NewQuizMarkDao(ctx)
	quizs, err := dao.GetAllQuiz(courseNumber)
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	for _, v := range quizs {
		quizMarks, err := dao4.GetQuizMarksByPerson(v.ID, userId)
		if err != nil {
			code = e.ERROR
			return serializar.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		var score float64
		for _, m := range quizMarks {
			score += m.Score
		}
		quizSum := &model.QuizSum{
			QuizId:       v.ID,
			QuizName:     v.QuizName,
			CourseNumber: v.CourseNumber,
			UserId:       userId,
			Score:        score,
			MaxScore:     v.MaxScore,
		}
		ans, _ := dao1.GetQuizSumByPerson(v.ID, userId)
		if ans.ID != 0 && ans != nil {
			_ = dao1.DeleteQuizSum(ans.ID)
		}
		err = dao1.CreateQuizSum(quizSum)
		if err != nil {
			code = e.ERROR
			return serializar.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		quizSums = append(quizSums, quizSum)
	}
	return serializar.Response{
		Status: code,
		Data:   serializar.BuildQuizSums(quizSums),
		Msg:    "enquiry success",
	}
}
