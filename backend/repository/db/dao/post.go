package dao

import (
	"9900project/repository/db/model"
	"context"

	"gorm.io/gorm"
)

type PostDao struct {
	*gorm.DB
}

func NewPostDao(ctx context.Context) *PostDao {
	return &PostDao{NewDBClient(ctx)}
}

func (dao *PostDao) CreatePost(post *model.Post) (err error) {
	err = dao.DB.Model(&model.Post{}).Create(&post).Error
	return err
}

func (dao *PostDao) UpdatePost(id uint, post *model.Post) (err error) {
	err = dao.DB.Model(&model.Post{}).Where("id=?", id).Updates(&post).Error
	return
}

func (dao *PostDao) GetPostById(id uint) (post *model.Post, err error) {
	err = dao.DB.Model(&model.Post{}).Where("id=?", id).First(&post).Error
	return
}

func (dao *PostDao) GetPostsByForumId(id uint) (posts []*model.Post, err error) {
	err = dao.DB.Model(&model.Post{}).Where("forum_id=?", id).Find(&posts).Error
	return
}

func (dao *PostDao) GetPostsByCourseNumber(courseNumber int) (posts []*model.Post, err error) {
	err = dao.DB.Model(&model.Post{}).Where("course_number=?", courseNumber).Find(&posts).Error
	return
}

func (dao *PostDao) DeletePost(id uint) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&model.Post{}).Error
	return
}

func (dao *PostDao) SearchPostsByInfo(info string, courseNumber int) (posts []*model.Post, err error) {
	if info == "" {
		return posts, nil
	}

	err = dao.DB.Model(&model.Post{}).Where("course_number=? AND (title LIKE ? or content LIKE ?)", courseNumber, "%"+info+"%", "%"+info+"%").Find(&posts).Error
	return
}
