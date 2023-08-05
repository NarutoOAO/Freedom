package dao

import (
	"9900project/repository/db/model"
	"context"

	"gorm.io/gorm"
)

type PostDao struct {
	*gorm.DB
}

// NewPostDao this file is to create some posts and update some posts in different courses discussion forum.
func NewPostDao(ctx context.Context) *PostDao {
	return &PostDao{NewDBClient(ctx)}
}

// CreatePost create post
func (dao *PostDao) CreatePost(post *model.Post) (err error) {
	err = dao.DB.Model(&model.Post{}).Create(&post).Error
	return err
}

// UpdatePost update post
func (dao *PostDao) UpdatePost(id uint, post *model.Post) (err error) {
	err = dao.DB.Model(&model.Post{}).Where("id=?", id).Updates(&post).Error
	return
}

func (dao *PostDao) GetPostById(id uint) (post *model.Post, err error) {
	err = dao.DB.Model(&model.Post{}).Where("id=?", id).First(&post).Error
	return
}

// GetPostsByForumId get posts by forum id
func (dao *PostDao) GetPostsByForumId(id uint) (posts []*model.Post, err error) {
	err = dao.DB.Model(&model.Post{}).Where("forum_id=?", id).Find(&posts).Error
	return
}

// GetPostsByCourseNumber get posts by course number
func (dao *PostDao) GetPostsByCourseNumber(courseNumber int) (posts []*model.Post, err error) {
	err = dao.DB.Model(&model.Post{}).Where("course_number=?", courseNumber).Find(&posts).Error
	return
}

// DeletePost delete post
func (dao *PostDao) DeletePost(id uint) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&model.Post{}).Error
	return
}

// SearchPostsByInfo search posts by info
func (dao *PostDao) SearchPostsByInfo(info string, courseNumber int) (posts []*model.Post, err error) {
	if info == "" {
		return posts, nil
	}
	err = dao.DB.Model(&model.Post{}).Where("course_number=? AND (title LIKE ? or content LIKE ?)", courseNumber, "%"+info+"%", "%"+info+"%").Find(&posts).Error
	return
}
