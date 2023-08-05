package serializar

import (
	"9900project/repository/db/model"
)

// create group
type TutorGroup struct {
	ID              uint   `json:"id"`
	CourseNumber    int    `json:"course_number"`
	GroupName       string `json:"group_name"`
	TeacherId       uint   `json:"teacher_id"`
	TeacherName     string `json:"teacher_name"`
	ResponsibleId   uint   `json:"responsible_id"`
	ResponsibleName string `json:"responsible_name"`
	AssignmentId    uint   `json:"assignment_id"`
	AssMarkId       uint   `json:"ass_mark_id"`
}

// build group
func BuildGroup(group *model.TutorGroup) *TutorGroup {
	return &TutorGroup{
		ID:              group.ID,
		CourseNumber:    group.CourseNumber,
		GroupName:       group.GroupName,
		TeacherId:       group.TeacherId,
		TeacherName:     group.TeacherName,
		ResponsibleId:   group.ResponsibleId,
		ResponsibleName: group.ResponsibleName,
		AssignmentId:    group.AssignmentId,
		AssMarkId:       group.AssMarkId,
	}
}

// build tutor groups
func BuildGroups(items []*model.TutorGroup) (groups []*TutorGroup) {
	for _, item := range items {
		group := BuildGroup(item)
		groups = append(groups, group)
	}
	return
}
