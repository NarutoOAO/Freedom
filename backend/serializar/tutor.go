package serializar

import (
	"9900project/repository/db/model"
)

// create tutor
type Tutor struct {
	ID           uint   `json:"id"`
	UserId       uint   `json:"user_id"`
	Email        string `json:"email"`
	NickName     string `json:"nick_name"`
	Authority    int    `json:"authority"`
	CourseNumber int    `json:"course_number"`
}

// build tutor
func BuildTutor(tutor *model.Tutor) *Tutor {
	return &Tutor{
		ID:           tutor.ID,
		UserId:       tutor.UserId,
		Email:        tutor.Email,
		NickName:     tutor.NickName,
		Authority:    tutor.Authority,
		CourseNumber: tutor.CourseNumber,
	}
}

// build tutors
func BuildTutors(items []*model.Tutor) (tutors []*Tutor) {
	for _, item := range items {
		tutor := BuildTutor(item)
		tutors = append(tutors, tutor)
	}
	return
}
