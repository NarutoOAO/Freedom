package dao

import (
	"context"
	"gorm.io/gorm"
)

type QuizDao struct {
	*gorm.DB
}

func NewQuizDao(ctx context.Context) *QuizDao {
	return &QuizDao{NewDBClient(ctx)}
}
