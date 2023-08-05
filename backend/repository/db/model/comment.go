package model

import "github.com/jinzhu/gorm"

// Comment this model for creating Comment
type Comment struct {
	gorm.Model
	PostId     uint
	AuthorId   uint
	AuthorName string
	Authority  int
	Content    string
}
