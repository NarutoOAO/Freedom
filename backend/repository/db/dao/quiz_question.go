package dao

import (
	"9900project/repository/db/model"
	"context"
	"gorm.io/gorm"
)

type QuizQuestionDao struct {
	*gorm.DB
}

// NewQuizQuestionDao this file is to set the quiz question
func NewQuizQuestionDao(ctx context.Context) *QuizQuestionDao {
	return &QuizQuestionDao{NewDBClient(ctx)}
}

func (dao *QuizQuestionDao) CreateQuizQuestion(quizQuestion *model.QuizQuestion) error {
	err := dao.DB.Model(&model.QuizQuestion{}).Create(&quizQuestion).Error
	return err
}

func (dao *QuizQuestionDao) GetQuizQuestionsByQuizId(quizId uint) (quizQuestions []*model.QuizQuestion, err error) {
	err = dao.DB.Model(&model.QuizQuestion{}).Where("quiz_id=?", quizId).Find(&quizQuestions).Error
	return
}

func (dao *QuizQuestionDao) GetQuizQuestionsById(quizId uint) (quizQuestion model.QuizQuestion, err error) {
	err = dao.DB.Model(&model.QuizQuestion{}).Where("id=?", quizId).First(&quizQuestion).Error
	return
}

func (dao *QuizQuestionDao) UpdateQuizQuestion(qId uint, quizQuestion *model.Quiz) error {
	err := dao.DB.Model(&model.QuizQuestion{}).Where("id=?", qId).Updates(&quizQuestion).Error
	return err
}
