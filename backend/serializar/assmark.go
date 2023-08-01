package serializar

import (
	"9900project/conf"
	"9900project/repository/db/model"
)

type AssMark struct {
	Id           uint    `json:"ass_mark_id"`
	GroupId      uint    `json:"group_id"`
	GroupName    string  `json:"group_name"`
	UserId       uint    `json:"user_id"`
	AssignmentId uint    `json:"assignment_id"`
	FileUrl      string  `json:"file_url"`
	Score        float64 `json:"score"`
	CourseNumber int     `json:"course_number"`
	MaxScore     float64 `json:"max_score"`
	Content      string  `json:"content"`
	AssName      string  `json:"ass_name"`
}

func BuildAssMark(assMark *model.AssMark) *AssMark {
	return &AssMark{
		Id:           assMark.ID,
		GroupId:      assMark.GroupId,
		GroupName:    assMark.GroupName,
		UserId:       assMark.StudentId,
		AssignmentId: assMark.AssignmentId,
		FileUrl:      conf.PhotoHost + conf.HttpPort + conf.AssSolutionPath + assMark.FileURL(),
		Score:        assMark.Mark,
		CourseNumber: assMark.CourseNumber,
		MaxScore:     assMark.MaxScore,
		Content:      assMark.Content,
		AssName:      assMark.AssName,
	}
}

func BuildAssMarks(items []*model.AssMark) (assMarks []*AssMark) {
	for _, item := range items {
		assMark := BuildAssMark(item)
		assMarks = append(assMarks, assMark)
	}
	return
}
