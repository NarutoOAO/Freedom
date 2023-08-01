package dao

import (
	"9900project/repository/db/model"
	"context"
	"gorm.io/gorm"
)

type ScoreDao struct {
	*gorm.DB
}

func NewScoreDao(ctx context.Context) *ScoreDao {
	return &ScoreDao{NewDBClient(ctx)}
}

func (dao *ScoreDao) CreateScore(score *model.Score) error {
	err := dao.DB.Model(&model.Score{}).Create(&score).Error
	return err
}

func (dao *ScoreDao) GetScoreByPerson(cN int, uId uint) (score *model.Score, err error) {
	err = dao.DB.Model(&model.Score{}).Where("course_number=? and user_id=?", cN, uId).First(&score).Error
	return
}

func (dao *ScoreDao) DeleteScore(qId uint) error {
	err := dao.DB.Where("id=?", qId).Delete(&model.Score{}).Error
	return err
}
