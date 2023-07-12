package dao

import (
	"9900project/repository/db/model"
	"context"
	"gorm.io/gorm"
)

type AssMarkDao struct {
	*gorm.DB
}

func NewAssMarkDao(ctx context.Context) *AssMarkDao {
	return &AssMarkDao{NewDBClient(ctx)}
}

func (dao *AssMarkDao) CreateAssMark(assMark *model.AssMark) error {
	err := dao.DB.Model(&model.AssMark{}).Create(&assMark).Error
	return err
}

func (dao *AssMarkDao) DeleteAssMark(aId uint) error {
	err := dao.DB.Where("id=?", aId).Delete(&model.AssMark{}).Error
	return err
}

func (dao *AssMarkDao) GetAssMark(id uint) (assMark *model.AssMark, err error) {
	err = dao.DB.Model(&model.AssMark{}).Where("id=?", id).First(&assMark).Error
	return
}

func (dao *AssMarkDao) GetAssMarkById(uId uint, courseNumber int) (assMarks []*model.AssMark, err error) {
	err = dao.DB.Model(&model.AssMark{}).Where("student_id=? and course_number=?", uId, courseNumber).Find(&assMarks).Error
	return
}

func (dao *AssMarkDao) GetAssMarkByCourseNumber(courseNumber int, aId uint) (assMarks []*model.AssMark, err error) {
	err = dao.DB.Model(&model.AssMark{}).Where("course_number=? and assignment_id=?", courseNumber, aId).Find(&assMarks).Error
	return
}

func (dao *AssMarkDao) GetAssMarkByAssId(uId uint, aId uint) (count int64, err error) {
	err = dao.DB.Model(&model.AssMark{}).Where("student_id=? and assignment_id=?", uId, aId).Count(&count).Error
	return
}

func (dao *AssMarkDao) UpdateAssMark(aId uint, assMark *model.AssMark) error {
	err := dao.DB.Model(&model.AssMark{}).Where("id=?", aId).Updates(&assMark).Error
	return err
}
