package v1

import (
	"9900project/pkg/e"
	util "9900project/pkg/utils"
	"9900project/serializar"
	service2 "9900project/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCourse(c *gin.Context) {
	var service *service2.CourseService
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if claim.Authority != 1 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "权限不足"})
		return
	}
	err := c.ShouldBindJSON(&service)
	if err == nil {
		res := service.CreateCourse(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		fmt.Printf("", err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": "参数绑定错误"})
	}
}

func GetCoursesById(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if claim.Authority == 1 {
		var service *service2.CourseService
		if err := c.ShouldBindJSON(service); err != nil {
			res := service.ShowTeacherCourseList(c.Request.Context(), claim.ID)
			c.JSON(http.StatusOK, res)
		} else {
			code := e.ERROR
			c.JSON(http.StatusBadRequest, serializar.Response{
				Status: code,
				Error:  e.GetMsg(code),
			})
			util.LogrusObj.Info(err)
		}
	} else if claim.Authority == 0 {
		var service *service2.CourseSelect
		if err := c.ShouldBindJSON(service); err != nil {
			res := service.GetCoursesSelectById(c.Request.Context(), claim.ID)
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
}

func GetCoursesByNumber(c *gin.Context) {
	var service *service2.CourseSelect
	err := c.ShouldBind(&service)
	if err == nil {
		res := service.GetCoursesByNumber(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		fmt.Printf("", err)
		code := e.ERROR
		c.JSON(http.StatusBadRequest, serializar.Response{
			Status: code,
			Error:  e.GetMsg(code),
		})
	}
}

func StudentSelectCourse(c *gin.Context) {
	var service *service2.CourseSelect
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if claim.Authority != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "It's students' function"})
		return
	}
	if err := c.ShouldBindJSON(service); err != nil {
		res := service.StudentSelectCourse(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		code := e.ERROR
		c.JSON(http.StatusBadRequest, serializar.Response{
			Status: code,
			Error:  e.GetMsg(code),
		})
	}
}
