package dao

import (
	"9900project/repository/db/model"
	"context"

	"gorm.io/gorm"
)

type MaterialDao struct {
	*gorm.DB
}

func NewMaterialDao(ctx context.Context) *MaterialDao {
	return &MaterialDao{NewDBClient(ctx)}
}

func (dao *MaterialDao) CreateMaterial(material *model.Material) error {
	err := dao.DB.Model(&model.Material{}).Create(&material).Error
	return err
}

func (dao *MaterialDao) GetMaterialById(mId uint) (material *model.Material, err error) {
	err = dao.DB.Model(&model.Material{}).Where("id=?", mId).First(&material).Error
	return
}

func (dao *MaterialDao) UpdateMaterial(mId uint, material *model.Material) error {
	err := dao.DB.Model(&model.Material{}).Where("id=?", mId).Updates(&material).Error
	return err
}

func (dao *MaterialDao) DeleteMaterial(mId uint) error {
	err := dao.DB.Where("id=?", mId).Delete(&model.Material{}).Error
	return err
}

func (dao *MaterialDao) GetAllMaterials(courseNumber int, fileCategory string) (materials []*model.Material, err error) {
	err = dao.DB.Model(&model.Material{}).Where("course_number=? and file_Category=?", courseNumber, fileCategory).Find(&materials).Error
	return
}

func (dao *MaterialDao) GetPartOfMaterials(courseNumber int, fileCategory string) (materials []*model.Material, err error) {
	err = dao.DB.Model(&model.Material{}).Where("course_number=? and file_Category=? and publish = 1", courseNumber, fileCategory).Find(&materials).Error
	return
}

func (dao *MaterialDao) GetAllMaterialsByInfo(courseNumber int, info string) (materials []*model.Material, err error) {
	if info == "" {
		return materials, nil
	}
	err = dao.DB.Model(&model.Material{}).Where("course_number=? and (file_Category LIKE ? or file_name LIKE ?)", courseNumber, "%"+info+"%", "%"+info+"%").Find(&materials).Error
	return
}

func (dao *MaterialDao) GetPartOfMaterialsByInfo(courseNumber int, info string) (materials []*model.Material, err error) {
	if info == "" {
		return materials, nil
	}
	err = dao.DB.Model(&model.Material{}).Where("course_number=? and publish = 1 and (file_Category LIKE ? or file_name LIKE ?)",
		courseNumber, "%"+info+"%", "%"+info+"%").Find(&materials).Error
	return
}
