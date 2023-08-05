package v1

import (
	util "9900project/pkg/utils"
	service2 "9900project/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// controller for group
func CreateGroup(c *gin.Context) {
	service := &service2.CreateGroupService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.CreateGroup(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "bind parameter error"})
	}
}

// get group list
func GetGroup(c *gin.Context) {
	service := &service2.GetGroupService{}
	courseNumber := c.Param("course_number")
	cN, _ := strconv.Atoi(courseNumber)
	response := service.GetGroups(c.Request.Context(), cN)
	c.JSON(http.StatusOK, response)
}

// delete group
func DeleteGroup(c *gin.Context) {
	service := &service2.GetGroupService{}
	id := c.Param("id")
	tId, _ := strconv.Atoi(id)
	response := service.DeleteGroupById(c.Request.Context(), uint(tId))
	c.JSON(http.StatusOK, response)
}

// update group info by id
func UpdateTutor(c *gin.Context) {
	service := &service2.GetGroupService{}
	id := c.Param("id")
	tId, _ := strconv.Atoi(id)
	if err := c.ShouldBind(&service); err == nil {
		res := service.UpdateGroupByTutor(c.Request.Context(), uint(tId))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "bind parameter error"})
	}
}

// get group info by id
func GetGroupByUserId(c *gin.Context) {
	service := &service2.GetGroupService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	courseNumber := c.Param("course_number")
	cN, _ := strconv.Atoi(courseNumber)
	response := service.GetGroupsByUserId(c.Request.Context(), cN, uint(claim.ID))
	c.JSON(http.StatusOK, response)
}
