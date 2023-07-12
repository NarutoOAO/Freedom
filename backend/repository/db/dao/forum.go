package dao

import (
	"9900project/repository/db/model"
	"context"
	"gorm.io/gorm"
)

type ForumDao struct {
	*gorm.DB
}

func NewForumDao(ctx context.Context) *ForumDao {
	return &ForumDao{NewDBClient(ctx)}
}

func (dao *ForumDao) CreateForum(forum *model.Forum) error {
	return dao.DB.Model(&model.Forum{}).Create(&forum).Error
}

func (dao *ForumDao) DeleteForum(id uint) error {
	return dao.DB.Where("id=?", id).Delete(&model.Forum{}).Error
}

func (dao *ForumDao) UpdateForum(id uint, forum *model.Forum) error {
	return dao.DB.Model(&model.Forum{}).Where("id=?", id).Updates(&forum).Error
}

func (dao *ForumDao) GetForumById(id uint) (forum *model.Forum, err error) {
	err = dao.DB.Model(&model.Forum{}).Where("id=?", id).First(&forum).Error
	return
}

func (dao *ForumDao) GetForumsByCourseNumber(courseNumber int) (forums []*model.Forum, err error) {
	err = dao.DB.Model(&model.Forum{}).Where("course_number=?", courseNumber).Find(&forums).Error
	return
}
