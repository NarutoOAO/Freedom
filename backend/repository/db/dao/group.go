package dao

import (
	"9900project/repository/db/model"
	"context"

	"gorm.io/gorm"
)

type GroupDao struct {
	*gorm.DB
}

// NewGroupDao This function of this file is to divide tutors into different groups and mark the assignment by different group
func NewGroupDao(ctx context.Context) *GroupDao {
	return &GroupDao{NewDBClient(ctx)}
}

// create group
func (dao *GroupDao) CreateGroup(group *model.TutorGroup) error {
	err := dao.DB.Model(&model.TutorGroup{}).Create(&group).Error
	return err
}

// get groups by course number
func (dao *GroupDao) GetGroups(id int) (groups []*model.TutorGroup, err error) {
	err = dao.DB.Model(&model.TutorGroup{}).Where("course_number=?", id).Find(&groups).Error
	return
}

// delete group by group id
func (dao *GroupDao) DeleteTutorById(id uint) error {
	err := dao.DB.Where("id=?", id).Delete(&model.TutorGroup{}).Error
	return err
}

// get group by group id
func (dao *GroupDao) GetGroupById(id uint) (group *model.TutorGroup, err error) {
	err = dao.DB.Model(&model.TutorGroup{}).Where("id=?", id).First(&group).Error
	return
}

// update group by group id
func (dao *GroupDao) UpdateGroupByTutor(id uint, tutor_id uint, name string) error {
	data := map[string]interface{}{
		"responsible_id":   tutor_id,
		"responsible_name": name,
	}
	return dao.DB.Model(&model.TutorGroup{}).Where("id = ?", id).Updates(data).Error
}

// get groups by user id
func (dao *GroupDao) GetGroupsByUserId(courseNumber int, id uint) (groups []*model.TutorGroup, err error) {
	err = dao.DB.Model(&model.TutorGroup{}).Where("course_number=? and responsible_id=?", courseNumber, id).Find(&groups).Error
	return
}
