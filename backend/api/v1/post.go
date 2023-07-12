package v1

import (
	"9900project/pkg/e"
	util "9900project/pkg/utils"
	"9900project/repository/cache"
	"9900project/serializar"
	service2 "9900project/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var service *service2.PostService
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		response := service.CreatePost(c.Request.Context(), claim.ID)
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

func GetPostByForumId(c *gin.Context) {
	var service *service2.PostService
	//id := c.Request.Header.Get("listingid")
	id := c.Param("id")
	fId, _ := strconv.Atoi(id)
	response := service.GetPostsByForumId(c.Request.Context(), uint(fId))
	c.JSON(http.StatusOK, response)
}

func GetPostByCourseNumber(c *gin.Context) {
	var service *service2.PostService
	//id := c.Request.Header.Get("listingid")
	courseNumber := c.Param("course")
	cN, _ := strconv.Atoi(courseNumber)
	response := service.GetPostsByCourseNumber(c.Request.Context(), cN)
	c.JSON(http.StatusOK, response)
}

func GetPostInformationByForumId(c *gin.Context) {
	var service *service2.PostService
	//id := c.Request.Header.Get("listingid")
	id := c.Param("id")
	pId, _ := strconv.Atoi(id)
	response := service.GetPostInformationByForumId(c.Request.Context(), uint(pId))
	c.JSON(http.StatusOK, response)
}

func GetPost2(c *gin.Context) {
	order, _ := c.GetQuery("order")
	pageStr, ok := c.GetQuery("page")
	if !ok {
		pageStr = "1"
	}
	pageNum, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		pageNum = 1
	}
	posts := cache.GetPost(order, pageNum)
	fmt.Println(len(posts))
	c.JSON(http.StatusOK, serializar.Response{
		Status: http.StatusOK,
		Msg:    "search success",
		Data:   posts,
		Error:  "",
	})
}

func PostVote(c *gin.Context) {
	var vote *service2.VoteData
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBindJSON(&vote); err == nil {
		res := vote.VoteToPost(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		code := e.ERROR
		c.JSON(http.StatusBadRequest, serializar.Response{
			Status: code,
			Error:  e.GetMsg(code),
		})
	}
}

func SearchPostByInfo(c *gin.Context) {
	var service *service2.PostService
	courseNumber := c.Param("course")
	cN, _ := strconv.Atoi(courseNumber)

	var searchPosts *service2.SearchPostService
	if err := c.ShouldBind(&searchPosts); err == nil {
		res := service.SearchPosts(c.Request.Context(), searchPosts.Info, cN)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
