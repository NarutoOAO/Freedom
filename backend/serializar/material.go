package serializar

import (
	"9900project/conf"
	"9900project/repository/db/model"
)

type Material struct {
	Id           uint   `json:"material_id"`
	CourseNumber int    `json:"course_number"`
	FileName     string `json:"file_name"`
	FileUrl      string `json:"file_url"`
	FileCategory string `json:"file_category"`
	Type         int    `json:"type"`
	Publish      int    `json:"publish"`
}

func BuildMaterial(material *model.Material) *Material {
	return &Material{
		Id:           material.ID,
		CourseNumber: material.CourseNumber,
		FileName:     material.FileName,
		FileUrl:      conf.PhotoHost + conf.HttpPort + conf.CoursePath + material.FileURL(),
		FileCategory: material.FileCategory,
		Type:         material.Type,
		Publish:      material.Publish,
	}
}

func BuildMaterials(items []*model.Material) (materials []*Material) {
	for _, item := range items {
		material := BuildMaterial(item)
		materials = append(materials, material)
	}
	return
}
