package service

import (
	"9900project/pkg/e"
	dao2 "9900project/repository/db/dao"
	"9900project/repository/db/model"
	"9900project/serializar"
	"context"
)

// CreateGroupService is a struct to create group
type CreateGroupService struct {
	CourseNumber    int    `json:"course_number"`
	GroupName       string `json:"group_name"`
	TeacherId       uint   `json:"teacher_id"`
	TeacherName     string `json:"teacher_name"`
	ResponsibleId   uint   `json:"responsible_id"`
	ResponsibleName string `json:"responsible_name"`
	AssignmentId    uint   `json:"assignment_id"`
	AssMarkId       uint   `json:"ass_mark_id"`
}

// define a service to get group
type GetGroupService struct {
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

// create group
func (service *CreateGroupService) CreateGroup(ctx context.Context) serializar.Response {
	code := e.SUCCESS
	var err error
	dao := dao2.NewGroupDao(ctx)
	tutorGroup := &model.TutorGroup{
		CourseNumber:    service.CourseNumber,
		GroupName:       service.GroupName,
		TeacherId:       service.TeacherId,
		TeacherName:     service.TeacherName,
		ResponsibleId:   service.ResponsibleId,
		ResponsibleName: service.ResponsibleName,
		AssignmentId:    service.AssignmentId,
		AssMarkId:       service.AssMarkId,
	}
	err = dao.CreateGroup(tutorGroup)
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializar.Response{
		Status: code,
		Data:   serializar.BuildGroup(tutorGroup),
		Msg:    "Add group success",
	}
}

// get groups
func (service *GetGroupService) GetGroups(ctx context.Context, courseNumber int) serializar.Response {
	code := e.SUCCESS
	var err error
	var groups []*model.TutorGroup
	dao := dao2.NewGroupDao(ctx)
	groups, err = dao.GetGroups(courseNumber)
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializar.Response{
		Status: code,
		Data:   serializar.BuildGroups(groups),
		Msg:    "enquiry success",
	}
}

// delete group by id
func (service *GetGroupService) DeleteGroupById(ctx context.Context, id uint) serializar.Response {
	code := e.SUCCESS
	var err error
	dao := dao2.NewGroupDao(ctx)
	group, _ := dao.GetGroupById(id)
	if group.ResponsibleId != 0 {
		return serializar.Response{
			Status: code,
			Msg:    "The group cannot be delete, there is a tutor in this group!",
		}
	}
	err = dao.DeleteTutorById(id)
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializar.Response{
		Status: code,
		Msg:    "delete success",
	}
}

// update group by tutor
func (service *GetGroupService) UpdateGroupByTutor(ctx context.Context, id uint) serializar.Response {
	code := e.SUCCESS
	dao := dao2.NewGroupDao(ctx)
	_, err := dao.GetGroupById(id)
	if err != nil {
		return serializar.Response{
			Status: code,
			Msg:    "Cannot find this group!",
		}
	}
	err1 := dao.UpdateGroupByTutor(id, service.ResponsibleId, service.ResponsibleName)
	if err1 != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err1.Error(),
		}
	}
	return serializar.Response{
		Status: code,
		Msg:    "Succeed to add tutor to group!",
	}
}

// get groups by user id
func (service *GetGroupService) GetGroupsByUserId(ctx context.Context, courseNumber int, id uint) serializar.Response {
	code := e.SUCCESS
	var err error
	var groups []*model.TutorGroup
	dao := dao2.NewGroupDao(ctx)
	groups, err = dao.GetGroupsByUserId(courseNumber, id)
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializar.Response{
		Status: code,
		Data:   serializar.BuildGroups(groups),
		Msg:    "enquiry success",
	}
}
