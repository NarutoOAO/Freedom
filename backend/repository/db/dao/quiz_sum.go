package dao

import (
	"9900project/repository/db/model"
	"context"
	"gorm.io/gorm"
)

type QuizSumDao struct {
	*gorm.DB
}

// NewQuizSumDao this file is to get sum of quiz mark.
func NewQuizSumDao(ctx context.Context) *QuizSumDao {
	return &QuizSumDao{NewDBClient(ctx)}
}

func (dao *QuizSumDao) CreateQuizSum(quiz *model.QuizSum) error {
	err := dao.DB.Model(&model.QuizSum{}).Create(&quiz).Error
	return err
}

func (dao *QuizSumDao) GetQuizSumById(uId uint) (quizSum *model.QuizSum, err error) {
	err = dao.DB.Model(&model.QuizSum{}).Where("user_id=?", uId).First(&quizSum).Error
	return
}

func (dao *QuizSumDao) GetQuizSumByPerson(qId uint, uId uint) (quizSum *model.QuizSum, err error) {
	err = dao.DB.Model(&model.QuizSum{}).Where("quiz_id=? and user_id=?", qId, uId).First(&quizSum).Error
	return
}

func (dao *QuizSumDao) UpdateQuizSumDao(qId uint, quizSum *model.QuizSum) error {
	err := dao.DB.Model(&model.QuizSum{}).Where("id=?", qId).Updates(&quizSum).Error
	return err
}

func (dao *QuizSumDao) DeleteQuizSum(qId uint) error {
	err := dao.DB.Where("id=?", qId).Delete(&model.QuizSum{}).Error
	return err
}

func (dao *QuizSumDao) GetAllQuizSum(courseNumber int) (quizSum []*model.QuizSum, err error) {
	err = dao.DB.Model(&model.QuizSum{}).Where("course_number=?", courseNumber).Find(&quizSum).Error
	return
}
