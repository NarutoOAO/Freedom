package cache

import (
	"github.com/go-redis/redis"
	"math"
	"time"
)

const (
	OneWeekInSeconds         = 7 * 24 * 3600
	VoteScore        float64 = 432
	PostPerAge               = 20
)

func PostVote(postID, userID string, v float64) (err error) {
	postTime := RedisClient.ZScore(KeyPostTimeZSet, postID).Val()
	if float64(time.Now().Unix())-postTime > OneWeekInSeconds {
		return ErrorVoteTimeExpire
	}
	key := KeyPostVotedZSetPrefix + postID
	ov := RedisClient.ZScore(key, userID).Val()

	diffAbs := math.Abs(ov - v)
	pipeline := RedisClient.TxPipeline()
	pipeline.ZAdd(key, redis.Z{
		Score:  v,
		Member: userID,
	})
	pipeline.ZIncrBy(KeyPostScoreZSet, VoteScore*diffAbs*v, postID)

	switch math.Abs(ov) - math.Abs(v) {
	case 1:
		pipeline.HIncrBy(KeyPostInfoHashPrefix+postID, "votes", -1)
	case 0:
	case -1:
		pipeline.HIncrBy(KeyPostInfoHashPrefix+postID, "votes", 1)
	default:
		return ErrorVoted
	}
	_, err = pipeline.Exec()
	return
}

func CreatePost(postID, userID, title, summary, ForumId string) (err error) {
	now := float64(time.Now().Unix())
	//KeyPostVotedZSetPrefix = "post:voted:"
	votedKey := KeyPostVotedZSetPrefix + postID
	//KeyForumPostSetPrefix = "forum:"
	communityKey := KeyForumPostSetPrefix + ForumId
	postInfo := map[string]interface{}{
		"title":    title,
		"summary":  summary,
		"post:id":  postID,
		"user:id":  userID,
		"time":     now,
		"votes":    1,
		"comments": 0,
	}

	// 事务操作
	pipeline := RedisClient.TxPipeline()
	pipeline.ZAdd(votedKey, redis.Z{
		Score:  1,
		Member: userID,
	})
	pipeline.Expire(votedKey, time.Second*OneWeekInSeconds)

	pipeline.HMSet(KeyPostInfoHashPrefix+postID, postInfo)
	pipeline.ZAdd(KeyPostScoreZSet, redis.Z{
		Score:  now + VoteScore,
		Member: postID,
	})
	pipeline.ZAdd(KeyPostTimeZSet, redis.Z{
		Score:  now,
		Member: postID,
	})
	pipeline.SAdd(communityKey, postID)
	_, err = pipeline.Exec()
	return
}

func GetPost(order string, page int64) []map[string]string {
	key := KeyPostScoreZSet
	if order == "time" {
		key = KeyPostTimeZSet
	}
	start := (page - 1) * PostPerAge
	end := start + PostPerAge - 1
	ids := RedisClient.ZRevRange(key, start, end).Val()
	postList := make([]map[string]string, 0, len(ids))
	for _, id := range ids {
		postData := RedisClient.HGetAll(KeyPostInfoHashPrefix + id).Val()
		postData["id"] = id
		postList = append(postList, postData)
	}
	return postList
}

func GetCommunityPost(communityName, orderKey string, page int64) []map[string]string {
	key := orderKey + communityName

	if RedisClient.Exists(key).Val() < 1 {
		RedisClient.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",
		}, KeyForumPostSetPrefix+communityName, orderKey)
		RedisClient.Expire(key, 60*time.Second)
	}
	return GetPost(key, page)
}
