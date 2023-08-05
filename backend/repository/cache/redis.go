package cache

import (
	"9900project/conf"
	"github.com/go-redis/redis"
	"strconv"
)

var RedisClient *redis.Client

func InitCache() {
	Redis()
}

func Redis() {
	db, _ := strconv.ParseUint(conf.RedisDbName, 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:     conf.RedisAddr,
		Password: conf.RedisPw,
		DB:       int(db),
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	RedisClient = client
}
