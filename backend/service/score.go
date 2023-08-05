package service

import (
	"9900project/pkg/e"
	dao2 "9900project/repository/db/dao"
	"9900project/repository/db/model"
	"9900project/serializar"
	"context"
)

type ScoreService struct {
}

func (service *ScoreService) GetScore(ctx context.Context, userId uint, courseNumber int) serializar.Response {
	code := e.SUCCESS
	var err error
	dao := dao2.NewQuizSumDao(ctx)
	dao1 := dao2.NewAssMarkDao(ctx)
	dao3 := dao2.NewScoreDao(ctx)
	var score float64
	var maxScore float64
	quizSums, err := dao.GetAllQuizSum(courseNumber)
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	for _, v := range quizSums {
		score += v.Score
		maxScore += v.MaxScore
	}
	assMarks, err := dao1.GetAssMarkById(userId, courseNumber)
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	for _, v := range assMarks {
		score += v.Mark
		maxScore += v.MaxScore
	}
	s := &model.Score{
		UserId:       userId,
		CourseNumber: courseNumber,
		Score:        score,
		MaxScore:     maxScore,
	}
	ans, _ := dao3.GetScoreByPerson(courseNumber, userId)
	if ans.ID != 0 && ans != nil {
		_ = dao3.DeleteScore(ans.ID)
	}
	err = dao3.CreateScore(s)
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
		Data:   serializar.BuildScore(s),
		Msg:    "enquiry success",
	}
}
