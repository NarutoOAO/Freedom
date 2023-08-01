package serializar

import "9900project/repository/db/model"

type QuizSum struct {
	QuizId       uint
	QuizName     string
	CourseNumber int
	UserId       uint
	Score        float64
	MaxScore     float64
}

func BuildQuizSum(quizSum *model.QuizSum) *QuizSum {
	return &QuizSum{
		QuizId:       quizSum.QuizId,
		QuizName:     quizSum.QuizName,
		CourseNumber: quizSum.CourseNumber,
		UserId:       quizSum.UserId,
		Score:        quizSum.Score,
		MaxScore:     quizSum.MaxScore,
	}
}

func BuildQuizSums(items []*model.QuizSum) (quizSums []*QuizSum) {
	for _, item := range items {
		quizSum := BuildQuizSum(item)
		quizSums = append(quizSums, quizSum)
	}
	return
}
