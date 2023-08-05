package v1

import (
	service2 "9900project/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// controller for tutor
func CreateTutor(c *gin.Context) {
	service := &service2.CreateTutorService{}
	fmt.Println(service)
	if err := c.ShouldBind(&service); err == nil {
		res := service.CreateTutor(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "bind parameter error"})
	}
}

// get tutor list
func GetTutor(c *gin.Context) {
	service := &service2.GetTutorService{}
	courseNumber := c.Param("course_number")
	cN, _ := strconv.Atoi(courseNumber)
	response := service.GetTutors(c.Request.Context(), cN)
	c.JSON(http.StatusOK, response)
}

// delete tutor
func DeleteTutor(c *gin.Context) {
	service := &service2.GetTutorService{}
	id := c.Param("id")
	tId, _ := strconv.Atoi(id)
	response := service.DeleteTutorById(c.Request.Context(), uint(tId))
	c.JSON(http.StatusOK, response)
}
