package service

import (
	"9900project/pkg/e"
	util "9900project/pkg/utils"
	dao2 "9900project/repository/db/dao"
	"9900project/repository/db/model"
	"9900project/serializar"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

// NotificationService is a struct to create notification
type NotificationService struct {
	Id                uint            `json:"id"`
	Content           string          `json:"content"`
	PostId            uint            `json:"post_id"`
	Title             string          `json:"title"`
	CommentAuthorName string          `json:"comment_author_name"`
	PostAuthorId      uint            `json:"post_author_id"`
	Status            int             `json:"status"`
	CommentTime       model.LocalTime `json:"comment_time"`
	PostAuthorName    string          `json:"post_author_name"`
	CourseTeacherId   uint            `json:"course_teacher_id"`
	CourseTeacherName string          `json:"course_teacher_name"`
	CourseNumber      int             `json:"course_number"`
}

// CreateNotificationService is a struct to update notification
type UpdateNotificationService struct {
	Status int `json:"status"`
}

// define a service for websocket
type WebSocketClient struct {
	conn  *websocket.Conn
	token string
}

// define a map to store websocket clients
var clients = make(map[*websocket.Conn]WebSocketClient)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	Subprotocols: []string{"token"},
}

// define a function to handle websocket
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("WebSocket upgrade error: ", err)
		return
	}
	defer conn.Close()

	// Extract token from the subprotocol
	token := strings.Split(r.Header.Get("Sec-WebSocket-Protocol"), ", ")[1]
	log.Println("Token:", token)

	client := WebSocketClient{
		conn:  conn,
		token: token,
	}
	clients[conn] = client

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			delete(clients, conn)
			break
		}
		log.Println("Received message from client: ", string(msg))
	}
}

// define a function to connect websocket
func ConnectWebSocket() {
	http.HandleFunc("/ws", handleWebSocket)
}

func HandleMessages(notifications []*model.Notification, id uint) {
	messageJSON, err := json.Marshal(notifications)
	if err != nil {
		log.Println("Error marshaling note to JSON: ", err)
		return
	}

	for client := range clients {
		claims, err := util.ParseToken(clients[client].token)
		if err != nil {
			log.Println("Error parsing token: ", err)
			continue
		}

		if claims.ID == id {
			err := clients[client].conn.WriteMessage(websocket.TextMessage, messageJSON)
			if err != nil {
				log.Println("WebSocket write error: ", err)
				clients[client].conn.Close()
				delete(clients, client)
			}
		}
	}
}

// define a function to get notifications by id
func (service *NotificationService) GetNotificationsById(ctx context.Context, id uint) serializar.Response {
	code := e.SUCCESS
	dao := dao2.NewNotificationDao(ctx)
	notifications, err := dao.GetNotificationsById(id)
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    "database error",
			Error:  err.Error(),
		}
	}
	return serializar.Response{
		Status: code,
		Msg:    "enquire success",
		Data:   serializar.BuildNotifications(notifications),
	}
}

// define a function to update notification status by id
func (service *UpdateNotificationService) UpdatetNotification(ctx context.Context, id uint, status int) serializar.Response {
	code := e.SUCCESS
	dao := dao2.NewNotificationDao(ctx)
	err := dao.UpdatetNotification(id, status)
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    "database error",
			Error:  err.Error(),
		}
	}
	return serializar.Response{
		Status: code,
		Msg:    "enquire success",
	}
}
