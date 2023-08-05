package serializar

import "9900project/repository/db/model"

// create post
type Post struct {
	ID           uint   `json:"id"`
	ForumID      uint   `json:"forum_id"`
	ForumName    string `json:"forum_name"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	AuthorId     uint   `json:"author_id"`
	AuthorName   string `json:"author_name"`
	CourseNumber int    `json:"course_number"`
	View         uint64 `json:"view"`
}

// build post
func BuildPost(post *model.Post) *Post {
	return &Post{
		ID:           post.ID,
		ForumID:      post.ForumID,
		ForumName:    post.ForumName,
		Title:        post.Title,
		Content:      post.Content,
		AuthorId:     post.AuthorId,
		AuthorName:   post.AuthorName,
		CourseNumber: post.CourseNumber,
		//View:         post.View(),
	}
}

// build posts
func BuildPosts(items []*model.Post) (posts []*Post) {
	for _, item := range items {
		post := BuildPost(item)
		posts = append(posts, post)
	}
	return
}

// create data list
type DataList struct {
	Item  interface{} `json:"item"`
	Total uint        `json:"total"`
}

// build response data list
func BuildListResponse(items interface{}, total uint) Response {
	return Response{
		Status: 200,
		Data: DataList{
			Item:  items,
			Total: total,
		},
	}
}
