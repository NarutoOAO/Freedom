package serializar

import "9900project/repository/db/model"

type Score struct {
	UserId       uint
	CourseNumber int
	Score        float64
	MaxScore     float64
}

func BuildScore(score *model.Score) *Score {
	return &Score{
		UserId:       score.UserId,
		CourseNumber: score.CourseNumber,
		Score:        score.Score,
		MaxScore:     score.MaxScore,
	}
}

func BuildScores(items []*model.Score) (scores []*Score) {
	for _, item := range items {
		score := BuildScore(item)
		scores = append(scores, score)
	}
	return
}
