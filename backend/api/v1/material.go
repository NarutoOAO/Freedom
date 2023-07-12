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

func CreateMaterial(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	if fileHeader == nil {
		code := e.ERROR
		c.JSON(http.StatusBadRequest, serializar.Response{
			Status: code,
			Error:  e.GetMsg(code),
		})
	}
	fileSize := fileHeader.Size
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if claim.Authority != 1 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "权限不足"})
		return
	}
	var service *service2.MaterialService
	if err := c.ShouldBind(&service); err == nil {
		res := service.UploadMaterial(c.Request.Context(), file, fileSize)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "bind parameter error"})
	}
}

func ShowMaterial(c *gin.Context) {
	service := &service2.MaterialService{}
	courseNumber := c.Param("course_number")
	fC := c.Param("file_category")
	cN, _ := strconv.Atoi(courseNumber)
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	response := service.GetMaterials(c.Request.Context(), cN, fC, claim.Authority)
	c.JSON(http.StatusOK, response)
}

func UpdateMaterial(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if claim.Authority != 1 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "权限不足"})
		return
	}
	var service *service2.MaterialService
	if err := c.ShouldBind(&service); err == nil {
		res := service.UpdateMaterial(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "bind parameter error"})
	}
}

func DeleteMaterial(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if claim.Authority != 1 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "权限不足"})
		return
	}
	var service *service2.MaterialService
	if err := c.ShouldBind(&service); err == nil {
		res := service.DeleteMaterial(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "bind parameter error"})
	}
}

func SearchMaterialByInfo(c *gin.Context) {
	var service *service2.MaterialService
	var search *service2.SearchMaterialService
	courseNumber := c.Param("course_number")
	cN, _ := strconv.Atoi(courseNumber)
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&search); err == nil {
		res := service.GetMaterialsByInfo(c.Request.Context(), cN, claim.Authority, search.Info)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
