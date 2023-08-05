package serializar

import "9900project/repository/db/model"

type QuizMark struct {
	QuizId          uint
	QuestionId      uint
	QuestionNumber  int
	UserId          uint
	QuizDescription string
	QuizAnswer      string
	Type            int
	UserAnswer      string
	Score           float64
	MaxScore        float64
}

func BuildQuizMark(quizMark *model.QuizMark) *QuizMark {
	return &QuizMark{
		QuizId:          quizMark.QuizId,
		QuestionId:      quizMark.QuestionId,
		QuestionNumber:  quizMark.QuestionNumber,
		UserId:          quizMark.UserId,
		QuizDescription: quizMark.QuizDescription,
		QuizAnswer:      quizMark.QuizAnswer,
		Type:            quizMark.Type,
		UserAnswer:      quizMark.UserAnswer,
		Score:           quizMark.Score,
		MaxScore:        quizMark.MaxScore,
	}
}

func BuildQuizMarks(items []*model.QuizMark) (quizMarks []*QuizMark) {
	for _, item := range items {
		quizMark := BuildQuizMark(item)
		quizMarks = append(quizMarks, quizMark)
	}
	return
}
