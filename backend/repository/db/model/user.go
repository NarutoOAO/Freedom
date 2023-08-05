package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User model
type User struct {
	gorm.Model
	Email          string `gorm:"unique"`
	PasswordDigest string
	NickName       string
	Authority      int
	Avatar         string `gorm:"size:1000"`
	Studyoption    string
}

const (
	PassWordCost        = 12       //password
	Active       string = "active" //active users
)

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword check your password
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}

// AvatarUrl set avatar
func (user *User) AvatarURL() string {
	signedGetURL := user.Avatar
	return signedGetURL
}
