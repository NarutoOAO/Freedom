package v1

import (
	util "9900project/pkg/utils"
	service2 "9900project/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetQuizSum(c *gin.Context) {
	service := &service2.QuizSumService{}
	courseNumber := c.Param("course_number")
	cN, _ := strconv.Atoi(courseNumber)
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	response := service.GetQuizSumByStudent(c.Request.Context(), claim.ID, cN)
	c.JSON(http.StatusOK, response)
}
