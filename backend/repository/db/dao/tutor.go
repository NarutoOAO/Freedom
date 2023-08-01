package dao

import (
	"9900project/repository/db/model"
	"context"

	"gorm.io/gorm"
)

type TutorDao struct {
	*gorm.DB
}

func NewTutorDao(ctx context.Context) *TutorDao {
	return &TutorDao{NewDBClient(ctx)}
}

func (dao *TutorDao) CreateTutor(tutor *model.Tutor) error {
	err := dao.DB.Model(&model.Tutor{}).Create(&tutor).Error
	return err
}

func (dao *TutorDao) GetTutors(id int) (tutors []*model.Tutor, err error) {
	err = dao.DB.Model(&model.Tutor{}).Where("course_number=?", id).Find(&tutors).Error
	return
}

func (dao *TutorDao) DeleteTutorById(id uint) error {
	err := dao.DB.Where("id=?", id).Delete(&model.Tutor{}).Error
	return err
}

func (dao *TutorDao) IfExistOrNot(id uint, couseNumber int) (tutor *model.Tutor, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.Tutor{}).Where("course_number=? and user_id=?", couseNumber, id).Count(&count).Error
	if err != nil {
		return nil, false, err
	}
	if count == 0 {
		return nil, false, nil
	}
	err = dao.DB.Model(&model.Tutor{}).Where("course_number=? and user_id=?", couseNumber, id).First(&tutor).Error
	if err != nil {
		return nil, true, err
	}
	return tutor, true, nil
}
