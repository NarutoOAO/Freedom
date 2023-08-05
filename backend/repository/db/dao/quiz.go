package dao

import (
	"9900project/repository/db/model"
	"context"
	"gorm.io/gorm"
)

type QuizDao struct {
	*gorm.DB
}

// NewQuizDao this file is to create quiz in different courses.
func NewQuizDao(ctx context.Context) *QuizDao {
	return &QuizDao{NewDBClient(ctx)}
}

func (dao *QuizDao) CreateQuiz(quiz *model.Quiz) error {
	err := dao.DB.Model(&model.Quiz{}).Create(&quiz).Error
	return err
}

func (dao *QuizDao) GetQuizById(qId uint) (quiz *model.Quiz, err error) {
	err = dao.DB.Model(&model.Quiz{}).Where("id=?", qId).First(&quiz).Error
	return
}

func (dao *QuizDao) UpdateQuiz(qId uint, quiz *model.Quiz) error {
	err := dao.DB.Model(&model.Quiz{}).Where("id=?", qId).Updates(&quiz).Error
	return err
}

func (dao *QuizDao) DeleteQuiz(qId uint) error {
	err := dao.DB.Where("id=?", qId).Delete(&model.Quiz{}).Error
	return err
}

func (dao *QuizDao) GetAllQuiz(courseNumber int) (quiz []*model.Quiz, err error) {
	err = dao.DB.Model(&model.Quiz{}).Where("course_number=?", courseNumber).Find(&quiz).Error
	return
}
