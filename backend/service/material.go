package service

import (
	"9900project/pkg/e"
	util "9900project/pkg/utils"
	dao2 "9900project/repository/db/dao"
	"9900project/repository/db/model"
	"9900project/serializar"
	"context"
	"mime/multipart"
)

type MaterialService struct {
	MaterialId   int    `form:"material_id" json:"material_id"`
	Type         int    `form:"type" json:"type"`
	CourseNumber int    `form:"course_number" json:"course_number"`
	FileName     string `form:"file_name" json:"file_name"`
	FileCategory string `form:"file_category" json:"file_category"`
	Publish      int    `form:"publish" json:"publish"`
}

type SearchMaterialService struct {
	Info string `json:"info"`
}

func (service *MaterialService) UploadMaterial(ctx context.Context, file multipart.File, fileHeader int64) serializar.Response {
	code := e.SUCCESS
	var err error
	dao := dao2.NewMaterialDao(ctx)
	material := &model.Material{
		CourseNumber: service.CourseNumber,
		FileName:     service.FileName,
		FileCategory: service.FileCategory,
		Type:         service.Type,
		Publish:      service.Publish,
	}
	var path string
	if service.Type == 0 {
		path, err = util.UploadMaterialToLocalStatic1(file, service.CourseNumber, service.FileName)
	} else {
		path, err = util.UploadMaterialToLocalStatic2(file, service.CourseNumber, service.FileName)
	}
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	material.FileUrl = path
	err = dao.CreateMaterial(material)
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
		Data:   serializar.BuildMaterial(material),
		Msg:    "insert success",
	}
}

func (service *MaterialService) UpdateMaterial(ctx context.Context) serializar.Response {
	code := e.SUCCESS
	var err error
	dao := dao2.NewMaterialDao(ctx)
	material, err := dao.GetMaterialById(uint(service.MaterialId))
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	material.Publish = service.Publish
	err = dao.UpdateMaterial(uint(service.MaterialId), material)
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
		Data:   serializar.BuildMaterial(material),
		Msg:    "update success",
	}
}

func (service *MaterialService) DeleteMaterial(ctx context.Context) serializar.Response {
	code := e.SUCCESS
	var err error
	dao := dao2.NewMaterialDao(ctx)
	err = dao.DeleteMaterial(uint(service.MaterialId))
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

func (service *MaterialService) GetMaterials(ctx context.Context, courseNumber int, fileCategory string, authority int) serializar.Response {
	code := e.SUCCESS
	var err error
	var materials []*model.Material
	dao := dao2.NewMaterialDao(ctx)
	if authority == 1 {
		materials, err = dao.GetAllMaterials(courseNumber, fileCategory)
	} else {
		materials, err = dao.GetPartOfMaterials(courseNumber, fileCategory)
	}
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
		Data:   serializar.BuildMaterials(materials),
		Msg:    "enquiry success",
	}
}

func (service *MaterialService) GetMaterialsByInfo(ctx context.Context, courseNumber int, authority int, info string) serializar.Response {
	code := e.SUCCESS
	var err error
	var materials []*model.Material
	dao := dao2.NewMaterialDao(ctx)
	if authority == 1 {
		materials, err = dao.GetAllMaterialsByInfo(courseNumber, info)
	} else {
		materials, err = dao.GetPartOfMaterialsByInfo(courseNumber, info)
	}
	if err != nil {
		code = e.ERROR
		return serializar.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	count := len(materials)
	return serializar.BuildListResponse(serializar.BuildMaterials(materials), uint(count))
}
