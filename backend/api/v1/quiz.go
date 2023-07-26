package v1

import (
	util "9900project/pkg/utils"
	service2 "9900project/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateQuiz(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if claim.Authority != 1 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "authorization is not enough"})
		return
	}
	var service *service2.QuizService
	if err := c.ShouldBind(&service); err == nil {
		res := service.CreateQuiz(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "bind parameter error"})
	}
}

func GetQuiz(c *gin.Context) {
	service := &service2.QuizService{}
	courseNumber := c.Param("course_number")
	cN, _ := strconv.Atoi(courseNumber)
	response := service.GetQuizs(c.Request.Context(), cN)
	c.JSON(http.StatusOK, response)
}
