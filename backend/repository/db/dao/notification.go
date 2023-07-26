package dao

import (
	"9900project/repository/db/model"
	"context"

	"gorm.io/gorm"
)

type NotificationDao struct {
	*gorm.DB
}

func NewNotificationDao(ctx context.Context) *NotificationDao {
	return &NotificationDao{NewDBClient(ctx)}
}

func (dao *NotificationDao) CreateNotification(notification *model.Notification) error {
	return dao.DB.Model(&model.Notification{}).Create(&notification).Error
}

func (dao *NotificationDao) GetNotificationsById(id uint) ([]*model.Notification, error) {
	var notifications []*model.Notification
	err := dao.DB.Model(&model.Notification{}).Where("post_author_id=? or course_teacher_id=?", id, id).Find(&notifications).Error
	return notifications, err
}

func (dao *NotificationDao) UpdatetNotification(id uint, status int) error {
	return dao.DB.Model(&model.Notification{}).Where("id=?", id).Update("status", status).Error
}
