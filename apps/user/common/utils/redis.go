package utils

import "github.com/redis/go-redis/v9"

var rdb *redis.Client

func InitRDB(dsn string, pass string) *redis.Client {
	rdb = redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: pass,
		DB:       0,
	})
	return rdb
}
