package service

import (
	"9900project/pkg/e"
	dao2 "9900project/repository/db/dao"
	"9900project/repository/db/model"
	"9900project/serializar"
	"context"
)

type CreateGroupMarkService struct {
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

type GetGroupMarkService struct {
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

func (service *CreateGroupMarkService) CreateGroup(ctx context.Context) serializar.Response {
	code := e.SUCCESS
	var err error
	dao := dao2.NewGroupMarkDao(ctx)
	tutorGroup := &model.GroupMark{
		CourseNumber:    service.CourseNumber,
		GroupId:         service.GroupId,
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
		Data:   serializar.BuildGroupMark(tutorGroup),
		Msg:    "Add group success",
	}
}

func (service *GetGroupMarkService) GetGroups(ctx context.Context, id uint) serializar.Response {
	code := e.SUCCESS
	var err error
	var groups []*model.GroupMark
	dao := dao2.NewGroupMarkDao(ctx)
	groups, err = dao.GetGroups(id)
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
		Data:   serializar.BuildGroupMarks(groups),
		Msg:    "enquiry success",
	}
}

func (service *GetGroupMarkService) DeleteGroupById(ctx context.Context, id uint) serializar.Response {
	code := e.SUCCESS
	var err error
	dao := dao2.NewGroupMarkDao(ctx)
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

func (service *GetGroupMarkService) UpdateGroupById(ctx context.Context, id uint) serializar.Response {
	code := e.SUCCESS
	dao := dao2.NewGroupMarkDao(ctx)
	_, err := dao.GetGroupById(id)
	if err != nil {
		return serializar.Response{
			Status: code,
			Msg:    "Cannot find this group!",
		}
	}
	err1 := dao.UpdateGroupById(id, service.AssMarkId, service.AssignmentId)
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
