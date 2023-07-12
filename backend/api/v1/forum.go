package v1

import (
	util "9900project/pkg/utils"
	service2 "9900project/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateForum(c *gin.Context) {
	var service *service2.ForumService
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if claim.Authority != 1 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "权限不足"})
		return
	}
	if err := c.ShouldBind(&service); err == nil {
		res := service.CreateForum(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "参数绑定错误"})
	}
}

func ShowForumList(c *gin.Context) {
	var service *service2.ForumService
	courseNumber := c.Param("course_number")
	cNumber, _ := strconv.Atoi(courseNumber)
	response := service.ShowForumList(c.Request.Context(), cNumber)
	c.JSON(http.StatusOK, response)
}
