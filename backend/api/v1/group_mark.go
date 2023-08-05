package v1

import (
	service2 "9900project/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// controller for group mark
func CreateGroupMark(c *gin.Context) {
	service := &service2.CreateGroupMarkService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.CreateGroup(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "bind parameter error"})
	}
}

// controller for group mark
func GetGroupMark(c *gin.Context) {
	service := &service2.GetGroupMarkService{}
	id := c.Param("group_id")
	gId, _ := strconv.Atoi(id)
	response := service.GetGroups(c.Request.Context(), uint(gId))
	c.JSON(http.StatusOK, response)
}

// controller for delete group mark
func DeleteGroupMark(c *gin.Context) {
	service := &service2.GetGroupMarkService{}
	id := c.Param("id")
	tId, _ := strconv.Atoi(id)
	response := service.DeleteGroupById(c.Request.Context(), uint(tId))
	c.JSON(http.StatusOK, response)
}

// controller for update group mark
func UpdateGroupMark(c *gin.Context) {
	service := &service2.GetGroupMarkService{}
	id := c.Param("id")
	tId, _ := strconv.Atoi(id)
	if err := c.ShouldBind(&service); err == nil {
		res := service.UpdateGroupById(c.Request.Context(), uint(tId))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "bind parameter error"})
	}
}
