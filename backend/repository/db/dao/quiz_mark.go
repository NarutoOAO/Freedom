package dao

import (
	"9900project/repository/db/model"
	"context"

	"gorm.io/gorm"
)

type QuizMarkDao struct {
	*gorm.DB
}

// NewQuizMarkDao this file is to mark quiz.
func NewQuizMarkDao(ctx context.Context) *QuizMarkDao {
	return &QuizMarkDao{NewDBClient(ctx)}
}

func (dao *QuizMarkDao) CreateQuizMark(quizMark *model.QuizMark) error {
	err := dao.DB.Model(&model.QuizMark{}).Create(&quizMark).Error
	return err
}

func (dao *QuizMarkDao) GetQuizMarkById(qId uint) (quizMark *model.QuizMark, err error) {
	err = dao.DB.Model(&model.QuizMark{}).Where("id=?", qId).First(&quizMark).Error
	return
}

func (dao *QuizMarkDao) UpdateQuizMark(qId uint, quizMark *model.QuizMark) error {
	err := dao.DB.Model(&model.QuizMark{}).Where("id=?", qId).Updates(&quizMark).Error
	return err
}

func (dao *QuizMarkDao) DeleteQuizMark(qId uint) error {
	err := dao.DB.Where("id=?", qId).Delete(&model.QuizMark{}).Error
	return err
}

func (dao *QuizMarkDao) GetQuizMarkByPerson(qId uint, uId uint) (quizMark *model.QuizMark, err error) {
	err = dao.DB.Model(&model.QuizMark{}).Where("question_id=? and user_id=?", qId, uId).First(&quizMark).Error
	return
}

func (dao *QuizMarkDao) GetQuizMarksByPerson(qId uint, uId uint) (quizMarks []*model.QuizMark, err error) {
	err = dao.DB.Model(&model.QuizMark{}).Where("quiz_id=? and user_id=?", qId, uId).Find(&quizMarks).Error
	return
}
