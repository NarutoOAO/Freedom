package dao

import (
	"9900project/repository/db/model"
	"context"
	"gorm.io/gorm"
)

type AssignmentDao struct {
	*gorm.DB
}

func NewAssignmentDao(ctx context.Context) *AssignmentDao {
	return &AssignmentDao{NewDBClient(ctx)}
}

func (dao *AssignmentDao) CreateAssignment(assignment *model.Assignment) error {
	err := dao.DB.Model(&model.Assignment{}).Create(&assignment).Error
	return err
}

func (dao *AssignmentDao) GetAssignmentById(aId uint) (assignment *model.Assignment, err error) {
	err = dao.DB.Model(&model.Assignment{}).Where("id=?", aId).First(&assignment).Error
	return
}

func (dao *AssignmentDao) UpdateAssignment(aId uint, assignment *model.Assignment) error {
	err := dao.DB.Model(&model.Assignment{}).Where("id=?", aId).Updates(&assignment).Error
	return err
}

func (dao *AssignmentDao) DeleteAssignment(aId uint) error {
	err := dao.DB.Where("id=?", aId).Delete(&model.Assignment{}).Error
	return err
}

func (dao *AssignmentDao) GetAllAssignments(courseNumber int) (assignments []*model.Assignment, err error) {
	err = dao.DB.Model(&model.Assignment{}).Where("course_number=?", courseNumber).Find(&assignments).Error
	return
}
