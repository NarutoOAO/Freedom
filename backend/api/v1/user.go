package v1

import (
	"9900project/pkg/e"
	util "9900project/pkg/utils"
	"9900project/serializar"
	service2 "9900project/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var service service2.UserService
	if err := c.ShouldBind(&service); err == nil {
		service.UserRegister(c.Request.Context())
		response := service.UserLogin(c.Request.Context())
		c.JSON(http.StatusOK, response)
	} else {
		code := e.ERROR
		c.JSON(http.StatusBadRequest, serializar.Response{
			Status: code,
			Error:  e.GetMsg(code),
		})
		util.LogrusObj.Info(err)
	}
}

func UserLogin(c *gin.Context) {
	var service service2.UserService
	if err := c.ShouldBind(&service); err == nil {
		response := service.UserLogin(c.Request.Context())
		c.JSON(http.StatusOK, response)
	} else {
		code := e.ERROR
		c.JSON(http.StatusBadRequest, serializar.Response{
			Status: code,
			Error:  e.GetMsg(code),
		})
		util.LogrusObj.Info(err)
	}
}

func UploadAvatar(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	if fileHeader == nil {
		code := e.ERROR
		c.JSON(http.StatusBadRequest, serializar.Response{
			Status: code,
			Error:  e.GetMsg(code),
		})
	}
	fileSize := fileHeader.Size
	service := service2.UserService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.UploadAvatar(c.Request.Context(), claim.ID, file, fileSize)
		c.JSON(http.StatusOK, res)
	} else {
		code := e.ERROR
		c.JSON(http.StatusBadRequest, serializar.Response{
			Status: code,
			Error:  e.GetMsg(code),
		})
		util.LogrusObj.Info(err)
	}
}

func UpdateUser(c *gin.Context) {
	updateUserService := service2.UserService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&updateUserService); err == nil {
		res := updateUserService.UpdateUser(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		code := e.ERROR
		c.JSON(http.StatusBadRequest, serializar.Response{
			Status: code,
			Error:  e.GetMsg(code),
		})
		util.LogrusObj.Info(err)
	}
}

func ChangePassword(c *gin.Context) {
	service := service2.UserService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.ChangePassword(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		code := e.ERROR
		c.JSON(http.StatusBadRequest, serializar.Response{
			Status: code,
			Error:  e.GetMsg(code),
		})
		util.LogrusObj.Info(err)
	}
}

func GetUserByName(c *gin.Context) {
	var service *service2.UserService
	search := service2.SearchUserService{}
	if err := c.ShouldBind(&search); err == nil {
		res := service.GetUserByInfo(c.Request.Context(), search.Info)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
