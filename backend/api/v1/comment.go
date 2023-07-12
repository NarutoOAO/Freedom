package v1

import (
	util "9900project/pkg/utils"
	service2 "9900project/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateComment(c *gin.Context) {
	var service *service2.CommentService
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.CreateComment(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "参数绑定错误"})
	}
}

func GetCommentByPostId(c *gin.Context) {
	var service *service2.CommentService
	//id := c.Request.Header.Get("listingid")
	id := c.Param("id")
	pId, _ := strconv.Atoi(id)
	response := service.GetCommentByPostId(c.Request.Context(), uint(pId))
	c.JSON(http.StatusOK, response)
}
