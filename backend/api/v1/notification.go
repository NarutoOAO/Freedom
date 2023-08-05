package v1

import (
	"9900project/pkg/e"
	util "9900project/pkg/utils"
	"9900project/serializar"
	service2 "9900project/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// controller for notification
func GetNotifications(c *gin.Context) {
	service := &service2.NotificationService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		response := service.GetNotificationsById(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, response)
	} else {
		code := e.ERROR
		c.JSON(http.StatusBadRequest, serializar.Response{
			Status: code,
			Error:  e.GetMsg(code),
		})
		util.LogrusObj.Info(err)
	}
}

// controller for update notification
func UpdatetNotification(c *gin.Context) {
	service := &service2.UpdateNotificationService{}
	id := c.Param("notification_id")
	nId, _ := strconv.Atoi(id)
	if err := c.ShouldBind(&service); err == nil {
		response := service.UpdatetNotification(c.Request.Context(), uint(nId), service.Status)
		c.JSON(http.StatusOK, response)
	} else {
		code := e.ERROR
		c.JSON(http.StatusBadRequest, serializar.Response{
			Status: code,
			Error:  e.GetMsg(code),
		})
		util.LogrusObj.Info(err)
	}
}
