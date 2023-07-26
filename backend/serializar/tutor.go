package serializar

import (
	"9900project/repository/db/model"
)

type Tutor struct {
	ID           uint   `json:"id"`
	Email        string `json:"email"`
	NickName     string `json:"nick_name"`
	Authority    int    `json:"authority"`
	CourseNumber int    `json:"course_number"`
}

func BuildTutor(tutor *model.Tutor) *Tutor {
	return &Tutor{
		ID:           tutor.ID,
		Email:        tutor.Email,
		NickName:     tutor.NickName,
		Authority:    tutor.Authority,
		CourseNumber: tutor.CourseNumber,
	}
}

func BuildTutors(items []*model.Tutor) (tutors []*Tutor) {
	for _, item := range items {
		tutor := BuildTutor(item)
		tutors = append(tutors, tutor)
	}
	return
}
