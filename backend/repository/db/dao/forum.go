package dao

import (
	"9900project/repository/db/model"
	"context"

	"gorm.io/gorm"
)

type ForumDao struct {
	*gorm.DB
}

// NewForumDao create discussion forum, delete discussion forum, update forum, get discussion form by course id
func NewForumDao(ctx context.Context) *ForumDao {
	return &ForumDao{NewDBClient(ctx)}
}

// CreateForum create discussion forum
func (dao *ForumDao) CreateForum(forum *model.Forum) error {
	return dao.DB.Model(&model.Forum{}).Create(&forum).Error
}

// DeleteForum delete discussion forum
func (dao *ForumDao) DeleteForum(id uint) error {
	return dao.DB.Where("id=?", id).Delete(&model.Forum{}).Error
}

// UpdateForum update discussion forum
func (dao *ForumDao) UpdateForum(id uint, forum *model.Forum) error {
	return dao.DB.Model(&model.Forum{}).Where("id=?", id).Updates(&forum).Error
}

// GetForumByCourseId get discussion forum by course id
func (dao *ForumDao) GetForumById(id uint) (forum *model.Forum, err error) {
	err = dao.DB.Model(&model.Forum{}).Where("id=?", id).First(&forum).Error
	return
}

// GetForumsByCourseNumber get discussion forums by course number
func (dao *ForumDao) GetForumsByCourseNumber(courseNumber int) (forums []*model.Forum, err error) {
	err = dao.DB.Model(&model.Forum{}).Where("course_number=?", courseNumber).Find(&forums).Error
	return
}
