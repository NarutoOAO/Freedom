package dao

import (
	"9900project/repository/db/model"
	"context"

	"gorm.io/gorm"
)

type GroupDao struct {
	*gorm.DB
}

func NewGroupDao(ctx context.Context) *GroupDao {
	return &GroupDao{NewDBClient(ctx)}
}

func (dao *GroupDao) CreateGroup(group *model.TutorGroup) error {
	err := dao.DB.Model(&model.TutorGroup{}).Create(&group).Error
	return err
}

func (dao *GroupDao) GetGroups(id int) (groups []*model.TutorGroup, err error) {
	err = dao.DB.Model(&model.TutorGroup{}).Where("course_number=?", id).Find(&groups).Error
	return
}

func (dao *GroupDao) DeleteTutorById(id uint) error {
	err := dao.DB.Where("id=?", id).Delete(&model.TutorGroup{}).Error
	return err
}

func (dao *GroupDao) GetGroupById(id uint) (group *model.TutorGroup, err error) {
	err = dao.DB.Model(&model.TutorGroup{}).Where("id=?", id).First(&group).Error
	return
}
