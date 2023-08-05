package v1

import (
	util "9900project/pkg/utils"
	service2 "9900project/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// controller for quiz mark
func CreateQuizMark(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	var service *service2.QuizMarkService
	if err := c.ShouldBind(&service); err == nil {
		res := service.CreateQuizMark(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "bind parameter error"})
	}
}

func GetQuizMark(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	var service *service2.QuizMarkService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetQuizMark(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "bind parameter error"})
	}
}

//func GiveQuizMark(c *gin.Context) {
//	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
//	if claim.Authority != 1 {
//		c.JSON(http.StatusBadRequest, gin.H{"msg": "authorization is not enough"})
//		return
//	}
//	var service *service2.QuizMarkService
//	if err := c.ShouldBind(&service); err == nil {
//		res := service.GetQuizMark(c.Request.Context(), claim.ID)
//		c.JSON(http.StatusOK, res)
//	} else {
//		c.JSON(http.StatusBadRequest, gin.H{"msg": "bind parameter error"})
//	}
//}
