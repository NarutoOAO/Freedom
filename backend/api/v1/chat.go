package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const apiKey = "a515dd8b35e9479ba2b8300ba0ecc24f"

type Perception struct {
	InputText `json:"inputText"`
}

type InputText struct {
	Text string `json:"text"`
}

type UserInfo struct {
	ApiKey string `json:"apiKey"`
	UserId string `json:"userId"`
}

type MessageRequest struct {
	ReqType    int `json:"reqType"`
	Perception `json:"perception"`
	UserInfo   `json:"userInfo"`
}

type Mes struct {
	Intent struct {
		AppKey       string `json:"appKey"`
		Code         int    `json:"code"`
		OperateState int    `json:"operateState"`
		Parameters   struct {
			Date string `json:"date"`
			City string `json:"city"`
		} `json:"parameters"`
	} `json:"intent"`
	Results []struct {
		GroupType  int    `json:"groupType"`
		ResultType string `json:"resultType"`
		Values     struct {
			Text string `json:"text"`
		} `json:"values"`
	} `json:"results"`
}

func Chat(c *gin.Context) {
	input := &InputText{}
	err := c.ShouldBindJSON(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bind json error"})
	}
	perception := Perception{*input}
	messageRequest := MessageRequest{ReqType: 0, Perception: perception, UserInfo: UserInfo{ApiKey: apiKey, UserId: "1"}}
	messageRequest.UserInfo.ApiKey = apiKey

	res, _ := json.Marshal(messageRequest)
	fmt.Println(string(res))
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://openapi.tuling123.com/openapi/api/v2", bytes.NewBuffer(res))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending API request:", err)
		return
	}
	defer resp.Body.Close()

	// 处理API响应
	var mes Mes
	json.NewDecoder(resp.Body).Decode(&mes)

	// 提取ChatGPT的回复
	reply := mes
	fmt.Println("ChatGPT's Reply:", reply)
}
