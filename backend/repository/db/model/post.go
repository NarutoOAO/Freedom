package model

import (
	"9900project/repository/cache"
	"github.com/jinzhu/gorm"
	"strconv"
)

type Post struct {
	gorm.Model
	ForumID      uint
	Title        string
	Content      string
	AuthorId     uint
	Status       int
	CourseNumber int
	ForumName    string
	AuthorName   string
}

// View get the visit times
func (post *Post) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.PostViewKey(post.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

// AddView see the visit times
func (post *Post) AddView() {
	// add post key
	cache.RedisClient.Incr(cache.PostViewKey(post.ID))
	// add rank key
	cache.RedisClient.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(post.ID)))
}
