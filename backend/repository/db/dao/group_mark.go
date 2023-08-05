package dao

import (
	"9900project/repository/db/model"
	"context"

	"gorm.io/gorm"
)

type GroupMarkDao struct {
	*gorm.DB
}

// NewGroupMarkDao , this file is to mark the assignment by different group
func NewGroupMarkDao(ctx context.Context) *GroupMarkDao {
	return &GroupMarkDao{NewDBClient(ctx)}
}

// CreateGroup create group
func (dao *GroupMarkDao) CreateGroup(group *model.GroupMark) error {
	err := dao.DB.Model(&model.GroupMark{}).Create(&group).Error
	return err
}

// GetGroups get groups
func (dao *GroupMarkDao) GetGroups(id uint) (groups []*model.GroupMark, err error) {
	err = dao.DB.Model(&model.GroupMark{}).Where("group_id=?", id).Find(&groups).Error
	return
}

// DeleteGroup delete group
func (dao *GroupMarkDao) DeleteTutorById(id uint) error {
	err := dao.DB.Where("id=?", id).Delete(&model.GroupMark{}).Error
	return err
}

// GetGroupById get group by id
func (dao *GroupMarkDao) GetGroupById(id uint) (group *model.GroupMark, err error) {
	err = dao.DB.Model(&model.GroupMark{}).Where("id=?", id).First(&group).Error
	return
}

// UpdateGroupById update group by id
func (dao *GroupMarkDao) UpdateGroupById(id uint, ass_mark_id uint, assignment_id uint) error {
	data := map[string]interface{}{
		"ass_mark_id":   ass_mark_id,
		"assignment_id": assignment_id,
	}
	return dao.DB.Model(&model.GroupMark{}).Where("id = ?", id).Updates(data).Error
}
