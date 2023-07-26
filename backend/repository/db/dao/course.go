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
	return course, err
}

func (dao *CourseDao) GetAllCourses() (Courses []*model.Course, err error) {
	err = dao.DB.Model(&model.Course{}).Find(&Courses).Error
	return
}

func (dao *CourseDao) StudentSelectCourse(id uint) (Courses []*model.Course, err error) {
	err = dao.DB.Model(&model.Course{}).Where("course_number NOT IN ( SELECT course_number FROM course_select WHERE student_id = ? and deleted_at is null )", id).Find(&Courses).Error
	return
}

func (dao *CourseDao) GetByNotSelected(classification string, cnumber []int) (Courses []*model.Course, err error) {
	err = dao.DB.Model(&model.Course{}).Where("classification = ? and course_number not in ?", classification, cnumber).Find(&Courses).Error
	return
}
