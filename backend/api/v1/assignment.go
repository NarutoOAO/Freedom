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

// CreateAssignment controller for create assignment
func CreateAssignment(c *gin.Context) {
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
	if claim.Authority != 1 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Authorization is not enough"})
		return
	}
	var service *service2.AssignmentService
	if err := c.ShouldBind(&service); err == nil {
		res := service.UploadAssignment(c.Request.Context(), file, fileSize)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "bind parameter error"})
	}
}

// UpdateAssignment controller for update assignment
func UpdateAssignment(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if claim.Authority != 1 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Authorization is not enough"})
		return
	}
	var service *service2.AssignmentService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.UpdateAssignment(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "bind parameter error"})
	}
}

// ShowAssignment controller for show assignment
func ShowAssignment(c *gin.Context) {
	service := &service2.AssignmentService{}
	courseNumber := c.Param("course_number")
	cN, _ := strconv.Atoi(courseNumber)
	response := service.GetAssignments(c.Request.Context(), cN)
	c.JSON(http.StatusOK, response)
}

// DeleteAssignment controller for delete assignment
func DeleteAssignment(c *gin.Context) {
	var service *service2.AssignmentService
	if err := c.ShouldBind(&service); err == nil {
		res := service.DeleteAssignment(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "bind parameter error"})
	}
}
