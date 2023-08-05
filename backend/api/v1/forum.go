package v1

import (
	util "9900project/pkg/utils"
	service2 "9900project/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// controller for forum
func CreateForum(c *gin.Context) {
	var service *service2.ForumService
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if claim.Authority != 1 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "you dont have enough authorization"})
		return
	}
	if err := c.ShouldBind(&service); err == nil {
		res := service.CreateForum(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "bind parameters error"})
	}
}

// controller for forum list
func ShowForumList(c *gin.Context) {
	var service *service2.ForumService
	courseNumber := c.Param("course_number")
	cNumber, _ := strconv.Atoi(courseNumber)
	response := service.ShowForumList(c.Request.Context(), cNumber)
	c.JSON(http.StatusOK, response)
}
