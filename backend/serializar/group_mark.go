package serializar

import (
	"9900project/repository/db/model"
)

// create group mark
type GroupMark struct {
	ID              uint   `json:"id"`
	GroupId         uint   `json:"group_id"`
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
func BuildGroupMark(group *model.GroupMark) *GroupMark {
	return &GroupMark{
		ID:              group.ID,
		GroupId:         group.GroupId,
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

// build groups
func BuildGroupMarks(items []*model.GroupMark) (groups []*GroupMark) {
	for _, item := range items {
		group := BuildGroupMark(item)
		groups = append(groups, group)
	}
	return
}
