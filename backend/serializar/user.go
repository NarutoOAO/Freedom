package serializar

import (
	"9900project/conf"
	"9900project/repository/db/model"
)

type User struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	NickName  string `json:"nickname"`
	Authority int    `json:"authority"`
	Avatar    string `gorm:"size:1000"`
}

func BuildUser(user *model.User) *User {
	return &User{
		ID:        user.ID,
		Email:     user.Email,
		NickName:  user.NickName,
		Authority: user.Authority,
		Avatar:    conf.PhotoHost + conf.HttpPort + conf.AvatarPath + user.AvatarURL(),
	}
}
