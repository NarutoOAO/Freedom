package dao

import (
	"9900project/repository/db/model"
	"context"

	"gorm.io/gorm"
)

type CommentDao struct {
	*gorm.DB
}

// NewCommentDao get comment,delete comment, update comment, get post id
func NewCommentDao(ctx context.Context) *CommentDao {
	return &CommentDao{NewDBClient(ctx)}
}

// CreateComment create comment
func (dao *CommentDao) CreateComment(comment *model.Comment) error {
	return dao.DB.Model(&model.Comment{}).Create(&comment).Error
}

// delete comment
func (dao *CommentDao) DeleteComment(id uint) error {
	return dao.DB.Where("id=?", id).Delete(&model.Comment{}).Error
}

// update comment
func (dao *CommentDao) UpdateComment(id uint, comment *model.Comment) error {
	return dao.DB.Model(&model.Comment{}).Where("id=?", id).Updates(&comment).Error
}

// get comment by id
func (dao *CommentDao) GetCommentByPostId(id uint) (comments []*model.Comment, err error) {
	err = dao.DB.Model(&model.Comment{}).Where("post_id=?", id).Find(&comments).Error
	return
}
