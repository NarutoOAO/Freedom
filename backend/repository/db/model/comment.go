package model

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	PostId     uint
	AuthorId   uint
	AuthorName string
	Authority  int
	Content    string
}
