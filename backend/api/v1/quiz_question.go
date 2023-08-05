package v1

import (
	util "9900project/pkg/utils"
	service2 "9900project/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// controller for quiz question
func CreateQuizQuestion(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if claim.Authority != 1 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "authorization is not enough"})
		return
	}
	var service *service2.QuizQuestionService
	if err := c.ShouldBind(&service); err == nil {
		res := service.CreateQuizQuestion(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "bind parameter error"})
	}
}

func GetQuizQuestions(c *gin.Context) {
	qId := c.Param("quiz_id")
	quizId, _ := strconv.Atoi(qId)
	var service *service2.QuizQuestionService
	res := service.GetQuizQuestions(c.Request.Context(), uint(quizId))
	c.JSON(http.StatusOK, res)
}
