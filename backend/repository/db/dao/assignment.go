package dao

import (
	"9900project/repository/db/model"
	"context"
	"gorm.io/gorm"
)

type AssignmentDao struct {
	*gorm.DB
}

// NewAssignmentDao new assignment
func NewAssignmentDao(ctx context.Context) *AssignmentDao {
	return &AssignmentDao{NewDBClient(ctx)}
}

// CreateAssignment create assignment
func (dao *AssignmentDao) CreateAssignment(assignment *model.Assignment) error {
	err := dao.DB.Model(&model.Assignment{}).Create(&assignment).Error
	return err
}

// GetAssignmentById get assignment
func (dao *AssignmentDao) GetAssignmentById(aId uint) (assignment *model.Assignment, err error) {
	err = dao.DB.Model(&model.Assignment{}).Where("id=?", aId).First(&assignment).Error
	return
}

// UpdateAssignment update assignment
func (dao *AssignmentDao) UpdateAssignment(aId uint, assignment *model.Assignment) error {
	err := dao.DB.Model(&model.Assignment{}).Where("id=?", aId).Updates(&assignment).Error
	return err
}

// DeleteAssignment delete assignment
func (dao *AssignmentDao) DeleteAssignment(aId uint) error {
	err := dao.DB.Where("id=?", aId).Delete(&model.Assignment{}).Error
	return err
}

// GetAllAssignments get assignment
func (dao *AssignmentDao) GetAllAssignments(courseNumber int) (assignments []*model.Assignment, err error) {
	err = dao.DB.Model(&model.Assignment{}).Where("course_number=?", courseNumber).Find(&assignments).Error
	return
}
