package v1

import (
	util "9900project/pkg/utils"
	service2 "9900project/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// controller for score
func GetScore(c *gin.Context) {
	service := &service2.ScoreService{}
	courseNumber := c.Param("course_number")
	cN, _ := strconv.Atoi(courseNumber)
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	response := service.GetScore(c.Request.Context(), claim.ID, cN)
	c.JSON(http.StatusOK, response)
}
