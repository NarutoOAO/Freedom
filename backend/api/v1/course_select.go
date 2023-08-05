package v1

import (
	"9900project/pkg/e"
	util "9900project/pkg/utils"
	"9900project/serializar"
	service2 "9900project/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// controller for course select
func SelectCourse(c *gin.Context) {
	var service *service2.CourseSelect
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if claim.Authority != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "teachers can not select course"})
		return
	}
	if err := c.ShouldBind(&service); err == nil {
		res := service.SelectCourse(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "bind parameters error"})
	}
}

// controller for course select
func GetCoursesSelectById(c *gin.Context) {
	var service *service2.CourseSelect
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if claim.Authority != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "It's students' function"})
		return
	}
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

// define the controller for course drop
func DropCourseById(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if claim.Authority != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Cannot drop courses"})
		return
	}
	var service *service2.CourseSelect
	courseNumber := c.Param("courseNumber")
	cN, _ := strconv.Atoi(courseNumber)
	if err := c.ShouldBindJSON(service); err != nil {
		res := service.StudentDropCourse(c.Request.Context(), cN, claim.ID)
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

func Statistics(c *gin.Context) {
	var service *service2.CourseSelect
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if claim.Authority != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "It's students' function"})
		return
	}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Statistics(c.Request.Context(), claim.ID)
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
