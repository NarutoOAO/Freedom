package dao

import (
	"9900project/repository/db/model"
	"context"

	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

func (dao *UserDao) CreateUser(user *model.User) error {
	return dao.DB.Model(&model.User{}).Create(&user).Error
}

func (dao *UserDao) DeleteUser(id uint) error {
	return dao.DB.Where("id=?", id).Delete(&model.User{}).Error
}

func (dao *UserDao) UpdateUser(id uint, user *model.User) error {
	return dao.DB.Model(&model.User{}).Where("id=?", id).Updates(&user).Error
}

func (dao *UserDao) GetUserById(id uint) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("id=?", id).First(&user).Error
	return
}

func (dao *UserDao) GetUsersByName(name string) (users []*model.User, err error) {
	if name == "" {
		return users, nil
	}
	err = dao.DB.Model(&model.User{}).Where("nick_name LIKE ?", "%"+name+"%").Find(&users).Error
	return
}

func (dao *UserDao) IfExistOrNot(email string) (user *model.User, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.User{}).Where("email=?", email).Count(&count).Error
	if err != nil {
		return nil, false, err
	}
	if count == 0 {
		return nil, false, nil
	}
	err = dao.DB.Model(&model.User{}).Where("email=?", email).First(&user).Error
	if err != nil {
		return nil, true, err
	}
	return user, true, nil
}
