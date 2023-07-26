package dao

import (
	"9900project/repository/db/model"
	"context"
	"fmt"

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

func (dao *CourseSelectDao) GetCourseByCourseNumber(courseNumber int) (coursesSelect []*model.CourseSelect, err error) {
	err = dao.DB.Model(&model.CourseSelect{}).Where("course_number=?", courseNumber).Find(&coursesSelect).Error
	return
}

func (dao *CourseSelectDao) DropCourseById(courseNumber int, id uint) error {
	fmt.Println("wozaizheli")
	return dao.DB.Where("course_number = ? AND student_id = ?", courseNumber, id).Delete(&model.CourseSelect{}).Error
}

func (dao *CourseSelectDao) GetCourseByClassification(uid uint, classification string) (coursesSelect []*model.CourseSelect, err error) {
	err = dao.DB.Model(&model.CourseSelect{}).Where("student_id = ? and course_number in (select course_number from course where classification = ? and deleted_at is null)", uid, classification).Find(&coursesSelect).Error
	return
}
