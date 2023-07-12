package dao

import (
	"9900project/repository/db/model"
	"context"
	"gorm.io/gorm"
)

type CourseDao struct {
	*gorm.DB
}

func NewCourseDao(ctx context.Context) *CourseDao {
	return &CourseDao{NewDBClient(ctx)}
}

func (dao *CourseDao) CreateCourse(course *model.Course) error {
	return dao.DB.Model(&model.Course{}).Create(&course).Error
}

func (dao *CourseDao) DeleteCourse(id uint) error {
	return dao.DB.Where("id=?", id).Delete(&model.Course{}).Error
}

func (dao *CourseDao) UpdateCourse(id uint, course *model.Course) error {
	return dao.DB.Model(&model.Course{}).Where("id=?", id).Updates(&course).Error
}

func (dao *CourseDao) GetCourseByTeacherId(id uint) (courses []*model.Course, err error) {
	//err = dao.DB.Model(&model.Course{}).Where("teacher_id=?", id).Distinct("course_number").Find(&courses).Error
	err = dao.DB.Model(&model.Course{}).Where("teacher_id=?", id).Find(&courses).Error
	return
}

func (dao *CourseDao) GetCourseByCourseNumber(num int) (course *model.Course, err error) {
	err = dao.DB.Model(&model.Course{}).Where("course_number=?", num).Find(&course).Error
	return
}

func (dao *CourseDao) GetAllCourses() (Courses []*model.Course, err error) {
	err = dao.DB.Model(&model.Course{}).Find(&Courses).Error
	return
}
