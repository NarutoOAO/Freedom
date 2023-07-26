package v1

import (
	service2 "9900project/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateGroup(c *gin.Context) {
	service := &service2.CreateGroupService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.CreateGroup(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "bind parameter error"})
	}
}

func GetGroup(c *gin.Context) {
	service := &service2.GetGroupService{}
	courseNumber := c.Param("course_number")
	cN, _ := strconv.Atoi(courseNumber)
	response := service.GetGroups(c.Request.Context(), cN)
	c.JSON(http.StatusOK, response)
}

func DeleteGroup(c *gin.Context) {
	service := &service2.GetGroupService{}
	id := c.Param("id")
	tId, _ := strconv.Atoi(id)
	response := service.DeleteGroupById(c.Request.Context(), uint(tId))
	c.JSON(http.StatusOK, response)
}
