package serializar

import "9900project/repository/db/model"

type Forum struct {
	ID           uint
	ForumName    string
	Introduction string
	CourseNumber int
}

func BuildForum(forum *model.Forum) *Forum {
	return &Forum{
		ID:           forum.ID,
		ForumName:    forum.ForumName,
		Introduction: forum.Introduction,
		CourseNumber: forum.CourseNumber,
	}
}

func BuildForums(items []*model.Forum) (forums []*Forum) {
	for _, item := range items {
		forum := BuildForum(item)
		forums = append(forums, forum)
	}
	return
}
