package v1

import (
	"9900project/pkg/e"
	util "9900project/pkg/utils"
	"9900project/serializar"
	service2 "9900project/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateAssMark(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	if fileHeader == nil {
		code := e.ERROR
		c.JSON(http.StatusBadRequest, serializar.Response{
			Status: code,
			Error:  e.GetMsg(code),
		})
	}
	fileSize := fileHeader.Size
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if claim.Authority != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Authorization is not enough"})
		return
	}
	var service *service2.AssMarkService
	if err := c.ShouldBind(&service); err == nil {
		res := service.UploadAssSolution(c.Request.Context(), file, fileSize, claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "bind parameter error"})
	}
}

func DeleteAssMark(c *gin.Context) {
	var service *service2.AssMarkService
	if err := c.ShouldBind(&service); err == nil {
		res := service.DeleteAssSolution(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "bind parameter error"})
	}
}

func ShowAssMark(c *gin.Context) {
	service := &service2.AssMarkService{}
	courseNumber := c.Param("course_number")
	assignmentId := c.Param("assignment_id")
	cN, _ := strconv.Atoi(courseNumber)
	aId, _ := strconv.Atoi(assignmentId)
	response := service.GetAssMarks(c.Request.Context(), cN, uint(aId))
	c.JSON(http.StatusOK, response)
}

func ShowAssMarkForStudent(c *gin.Context) {
	service := &service2.AssMarkService{}
	courseNumber := c.Param("course_number")
	cN, _ := strconv.Atoi(courseNumber)
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	response := service.GetAssMarksByStudent(c.Request.Context(), claim.ID, cN)
	c.JSON(http.StatusOK, response)
}

func UpdateAssMark(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if claim.Authority != 1 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Authorization is not enough"})
		return
	}
	var service *service2.AssMarkService
	if err := c.ShouldBind(&service); err == nil {
		res := service.UpdateAssMark(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "bind parameter error"})
	}
}
