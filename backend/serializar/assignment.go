package serializar

import (
	"9900project/conf"
	"9900project/repository/db/model"
)

type Assignment struct {
	Id           uint            `json:"assignment_id"`
	CourseNumber int             `json:"course_number"`
	FileName     string          `json:"file_name"`
	FileUrl      string          `json:"file_url"`
	EndTime      model.LocalTime `json:"end_time"`
	MaxScore     float64         `json:"max_score"`
}

func BuildAssignment(assignment *model.Assignment) *Assignment {
	return &Assignment{
		Id:           assignment.ID,
		CourseNumber: assignment.CourseNumber,
		FileName:     assignment.FileName,
		FileUrl:      conf.PhotoHost + conf.HttpPort + conf.AssignmentPath + assignment.FileURL(),
		EndTime:      assignment.EndTime,
		MaxScore:     assignment.MaxScore,
	}
}

func BuildAssignments(items []*model.Assignment) (assignments []*Assignment) {
	for _, item := range items {
		assignment := BuildAssignment(item)
		assignments = append(assignments, assignment)
	}
	return
}
