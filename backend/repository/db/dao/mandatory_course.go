package dao

import (
	"9900project/repository/db/model"
	"context"
	"gorm.io/gorm"
)

type MandatoryCourseDao struct {
	*gorm.DB
}

func NewMandatoryCourseDao(ctx context.Context) *MandatoryCourseDao {
	return &MandatoryCourseDao{NewDBClient(ctx)}
}

func (dao *MandatoryCourseDao) GetByClassification(classification string) (courses []*model.MandatoryCourse, err error) {
	err = dao.DB.Model(&model.MandatoryCourse{}).Where("classification like concat('%',?,'%')", classification).Find(&courses).Error
	return
}
