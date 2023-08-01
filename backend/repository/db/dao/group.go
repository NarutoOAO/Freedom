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

func (dao *GroupDao) UpdateGroupByTutor(id uint, tutor_id uint, name string) error {
	data := map[string]interface{}{
		"responsible_id":   tutor_id,
		"responsible_name": name,
	}
	return dao.DB.Model(&model.TutorGroup{}).Where("id = ?", id).Updates(data).Error
}

func (dao *GroupDao) GetGroupsByUserId(courseNumber int, id uint) (groups []*model.TutorGroup, err error) {
	err = dao.DB.Model(&model.TutorGroup{}).Where("course_number=? and responsible_id=?", courseNumber, id).Find(&groups).Error
	return
}
