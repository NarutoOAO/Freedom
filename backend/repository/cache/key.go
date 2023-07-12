package cache

import (
	"fmt"
	"strconv"
)

const (
	RankKey                = "rank"
	KeyForumPostSetPrefix  = "forum:"
	KeyPostInfoHashPrefix  = "post:"
	KeyPostVotedZSetPrefix = "post:voted:"
	KeyPostTimeZSet        = "post:voted:time"
	KeyPostScoreZSet       = "post:voted:score"
)

func PostViewKey(id uint) string {
	return fmt.Sprintf("view:post:%s", strconv.Itoa(int(id)))
}
