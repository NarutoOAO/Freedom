package serializar

import "9900project/repository/db/model"

// create comment
type Comment struct {
	ID         uint   `json:"id"`
	Content    string `json:"content"`
	AuthorId   uint   `json:"author_id"`
	AuthorName string `json:"author_name"`
	Authority  int    `json:"authorization"`
}

// build comment
func BuildComment(comment *model.Comment) *Comment {
	return &Comment{
		ID:         comment.ID,
		Content:    comment.Content,
		AuthorId:   comment.AuthorId,
		AuthorName: comment.AuthorName,
		Authority:  comment.Authority,
	}
}

// build comments
func BuildComments(items []*model.Comment) (comments []*Comment) {
	for _, item := range items {
		comment := BuildComment(item)
		comments = append(comments, comment)
	}
	return
}
