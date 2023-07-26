package serializar

import "9900project/repository/db/model"

type Notification struct {
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

func BuildNotification(notification *model.Notification) *Notification {
	return &Notification{
		Id:                notification.ID,
		Content:           notification.Content,
		PostId:            notification.PostId,
		Title:             notification.Title,
		CommentAuthorName: notification.CommentAuthorName,
		PostAuthorId:      notification.PostAuthorId,
		Status:            notification.Status,
		CommentTime:       model.LocalTime(notification.CreatedAt),
		PostAuthorName:    notification.PostAuthorName,
		CourseTeacherId:   notification.CourseTeacherId,
		CourseTeacherName: notification.CourseTeacherName,
		CourseNumber:      notification.CourseNumber,
	}
}

func BuildNotifications(items []*model.Notification) (notifications []*Notification) {
	for _, item := range items {
		notification := BuildNotification(item)
		notifications = append(notifications, notification)
	}
	return
}
