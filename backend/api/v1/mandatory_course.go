package v1

import (
	util "9900project/pkg/utils"
	service2 "9900project/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMandatoryCourse(c *gin.Context) {
	var service *service2.MandatoryCourseService
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if claim.Authority != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "can not see this module"})
		return
	}
	res := service.GetUserMandatoryCourse(c.Request.Context(), claim.ID)
	c.JSON(http.StatusOK, res)
}
