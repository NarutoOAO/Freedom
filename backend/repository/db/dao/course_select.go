package dao

import (
	"9900project/repository/db/model"
	"context"
	"gorm.io/gorm"
)

type CourseSelectDao struct {
	*gorm.DB
}

func NewCourseSelectDao(ctx context.Context) *CourseSelectDao {
	return &CourseSelectDao{NewDBClient(ctx)}
}

func (dao *CourseSelectDao) CreateCourseSelect(courseSelect *model.CourseSelect) error {
	return dao.DB.Model(&model.CourseSelect{}).Create(&courseSelect).Error
}

func (dao *CourseSelectDao) GetCourseByStudentId(id uint) (coursesSelect []*model.CourseSelect, err error) {
	err = dao.DB.Model(&model.CourseSelect{}).Where("student_id=?", id).Find(&coursesSelect).Error
	return
}
