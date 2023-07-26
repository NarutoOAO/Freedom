package service

import (
	"9900project/pkg/e"
	util "9900project/pkg/utils"
	dao2 "9900project/repository/db/dao"
	"9900project/repository/db/model"
	"9900project/serializar"
	"context"
	"mime/multipart"
)

type UserService struct {
	Nickname           string `json:"nickname"`
	Email              string `json:"email"`
	Password           string `json:"password"`
	Authority          int    `json:"authority"`
	NewPassword        string `json:"new_password"`
	ConfirmNewPassword string `json:"confirm_new_password"`
	Studyoption        string `json:"studyoption"`
}

func (service *UserService) UserRegister(ctx context.Context) serializar.Response {
	code := e.SUCCESS
	var user *model.User
	var err error
	dao := dao2.NewUserDao(ctx)
	_, exist, err := dao.IfExistOrNot(service.Email)
	if exist {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    "user is existed",
		}
	}
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user = &model.User{
		Email:       service.Email,
		NickName:    service.Nickname,
		Authority:   service.Authority,
		Avatar:      "avatar.JPG",
		Studyoption: service.Studyoption,
	}
	if service.Password == "" {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    "password cannot be empty",
			Error:  err.Error(),
		}
	}
	err = user.SetPassword(service.Password)
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    "set password failed",
			Error:  err.Error(),
		}
	}
	err = dao.CreateUser(user)
	if err != nil {
		code = e.ErrorDatabase
		return serializar.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializar.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializar.BuildUser(user),
	}
}

func (service *UserService) UserLogin(ctx context.Context) serializar.Response {
	var token string
	var user *model.User
	var err error
	code := e.SUCCESS
	dao := dao2.NewUserDao(ctx)
	user, exist, err := dao.IfExistOrNot(service.Email)
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    "enquire failed",
		}
	}
	if !exist {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    "user is not existed",
		}
	}
	if !user.CheckPassword(service.Password) {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    "password is wrong",
		}
	}
	token, err = util.GenerateToken(user.ID, user.Email, user.Authority)
	return serializar.Response{
		Status: code,
		Data: serializar.TokenResponse{
			Token: token,
			User:  serializar.BuildUser(user),
		},
		Msg: "login success",
	}
}

func (service *UserService) UploadAvatar(ctx context.Context, userId uint, file multipart.File, fileHeader int64) serializar.Response {
	code := e.SUCCESS
	var user *model.User
	var err error
	dao := dao2.NewUserDao(ctx)
	user, err = dao.GetUserById(userId)
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	path, err := util.UploadAvatarToLocalStatic(file, userId, user.NickName)
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	user.Avatar = path
	err = dao.UpdateUser(userId, user)
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializar.Response{
		Status: code,
		Data:   serializar.BuildUser(user),
		Msg:    "更新成功",
	}
}

func (service *UserService) UpdateUser(ctx context.Context, id uint) serializar.Response {
	code := e.SUCCESS
	var user *model.User
	var err error
	userDao := dao2.NewUserDao(ctx)
	user, err = userDao.GetUserById(id)
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	if service.Nickname != "" {
		user.NickName = service.Nickname
	}
	err = userDao.UpdateUser(id, user)
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializar.Response{
		Status: code,
		Data:   serializar.BuildUser(user),
		Msg:    "更新成功",
	}
}

func (service *UserService) ChangePassword(ctx context.Context, id uint) serializar.Response {
	var token string
	var user *model.User
	var err error
	code := e.SUCCESS
	dao := dao2.NewUserDao(ctx)
	user, err = dao.GetUserById(id)
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    "user is not existed",
		}
	}
	if !user.CheckPassword(service.Password) {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    "password is wrong",
		}
	}
	if service.NewPassword == service.Password {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    "new password is same as old password",
		}
	}
	if service.NewPassword != service.ConfirmNewPassword {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    "new password is different as confirm-new-password",
		}
	}
	_ = user.SetPassword(service.NewPassword)
	err = dao.UpdateUser(id, user)
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    "set password failed",
			Error:  err.Error(),
		}
	}
	return serializar.Response{
		Status: code,
		Data: serializar.TokenResponse{
			Token: token,
			User:  serializar.BuildUser(user),
		},
		Msg: "change password success",
	}
}
