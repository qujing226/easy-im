package utils

import (
	"fmt"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func InitRDB(dsn string, pass string) *redis.Client {
	//dsn = "118.178.120.11:16379"
	//pass = "easy-chat"
	fmt.Println("init redis", dsn, pass)
	rdb = redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: pass,
		DB:       0,
	})
	return rdb
}
